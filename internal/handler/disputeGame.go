package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"golang.org/x/time/rate"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/pkg/contract"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"gorm.io/gorm"
)

type RetryDisputeGameClient struct {
	Client *contract.RateAndRetryDisputeGameClient
	DB     *gorm.DB
}

func NewRetryDisputeGameClient(db *gorm.DB, address common.Address, rpc *ethclient.Client, limit rate.Limit,
	burst int) (*RetryDisputeGameClient, error) {
	newDisputeGame, err := contract.NewDisputeGame(address, rpc)
	if err != nil {
		return nil, err
	}
	retryLimitGame := contract.NewRateAndRetryDisputeGameClient(newDisputeGame, limit, burst)
	return &RetryDisputeGameClient{Client: retryLimitGame, DB: db}, nil
}

func (r *RetryDisputeGameClient) ProcessDisputeGameCreated(ctx context.Context, evt *schema.SyncEvent) error {
	dispute := event.DisputeGameCreated{}
	err := dispute.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[processDisputeGameCreated] event data to DisputeGameCreated err: %s", err)
	}
	err = r.addDisputeGame(ctx, evt)
	if err != nil {
		return fmt.Errorf("[processDisputeGameCreated] addDisputeGame err: %s", err)
	}
	blockchain.AddContract(dispute.DisputeProxy)
	return nil
}

func (r *RetryDisputeGameClient) ProcessDisputeGameMove(ctx context.Context, evt *schema.SyncEvent) error {
	disputeGameMove := event.DisputeGameMove{}
	err := disputeGameMove.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] event data to disputeGameMove err: %s", err)
	}
	index := disputeGameMove.ParentIndex.Add(disputeGameMove.ParentIndex, big.NewInt(1))
	data, err := r.Client.RetryClaimData(ctx, &bind.CallOpts{}, index)
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] contract: %s, index: %d move event get claim data err: %s", evt.ContractAddress, index, err)
	}

	claimData := &schema.GameClaimData{
		GameContract: evt.ContractAddress,
		DataIndex:    index.Int64(),
		ParentIndex:  data.ParentIndex,
		CounteredBy:  data.CounteredBy.Hex(),
		Claimant:     data.Claimant.Hex(),
		Bond:         data.Bond.Uint64(),
		Claim:        hex.EncodeToString(data.Claim[:]),
		Position:     data.Position.Uint64(),
		Clock:        data.Clock.Int64(),
	}
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(claimData).Error
		if err != nil {
			return fmt.Errorf("[processDisputeGameMove] save dispute game err: %s\n ", err)
		}
		evt.Status = schema.EventValid
		err = tx.Save(evt).Error
		if err != nil {
			return fmt.Errorf("[processDisputeGameMove] save event err: %s\n ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RetryDisputeGameClient) ProcessDisputeGameResolve(ctx context.Context, evt *schema.SyncEvent) error {
	disputeResolved := event.DisputeGameResolved{}
	err := disputeResolved.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[processDisputeGameResolve] event data to disputeResolved err: %s", err)
	}
	disputeGame := &schema.DisputeGame{}
	err = r.DB.Where("game_contract=?", evt.ContractAddress).First(disputeGame).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("[processDisputeGameResolve] resolve event can not find dispute game err: %s", err)
	}
	disputeGame.Status = disputeResolved.Status
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(disputeGame).Error
		if err != nil {
			return fmt.Errorf("[processDisputeGameMove] update dispute game status err: %s\n ", err)
		}
		evt.Status = schema.EventValid
		err = tx.Save(evt).Error
		if err != nil {
			return fmt.Errorf("[processDisputeGameMove] update event err: %s\n ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	blockchain.RemoveContract(evt.ContractAddress)
	log.Infof("remove contract: %s", evt.ContractAddress)
	return nil
}

func (r *RetryDisputeGameClient) addDisputeGame(ctx context.Context, evt *schema.SyncEvent) error {
	disputeGame := event.DisputeGameCreated{}
	err := disputeGame.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[addDisputeGame] event data to DisputeGameCreated err: %s", err)
	}

	l2Block, err := r.Client.RetryL2BlockNumber(ctx, &bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET game L2BlockNumber err: %s", err)
	}
	status, err := r.Client.RetryStatus(ctx, &bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET game status err: %s", err)
	}
	claimData, err := r.Client.RetryClaimData(ctx, &bind.CallOpts{}, big.NewInt(0))
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET index 0 ClaimData err: %s", err)
	}

	gameClaim := &schema.GameClaimData{
		GameContract: strings.ToLower(disputeGame.DisputeProxy),
		DataIndex:    0,
		ParentIndex:  claimData.ParentIndex,
		CounteredBy:  claimData.CounteredBy.Hex(),
		Claimant:     claimData.Claimant.Hex(),
		Bond:         claimData.Bond.Uint64(),
		Claim:        hex.EncodeToString(claimData.Claim[:]),
		Position:     claimData.Position.Uint64(),
		Clock:        claimData.Clock.Int64(),
	}

	game := &schema.DisputeGame{
		SyncBlockID:     evt.SyncBlockID,
		Blockchain:      evt.Blockchain,
		BlockTime:       evt.BlockTime,
		BlockNumber:     evt.BlockNumber,
		BlockHash:       evt.BlockHash,
		BlockLogIndexed: evt.BlockLogIndexed,
		TxIndex:         evt.TxIndex,
		TxHash:          evt.TxHash,
		EventName:       evt.EventName,
		EventHash:       evt.EventHash,
		ContractAddress: evt.ContractAddress,
		GameContract:    strings.ToLower(disputeGame.DisputeProxy),
		GameType:        disputeGame.GameType,
		L2BlockNumber:   l2Block.Int64(),
		Status:          status,
		InitStatus:      schema.DisputeGameInit,
	}
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(gameClaim).Error
		if err != nil {
			return fmt.Errorf("[addDisputeGame] save dispute game claim: %s", err)
		}
		err = tx.Save(game).Error
		if err != nil {
			return fmt.Errorf("[addDisputeGame] save dispute game err: %s ", err)
		}
		evt.Status = schema.EventValid
		err = tx.Save(evt).Error
		if err != nil {
			return fmt.Errorf("[addDisputeGame] save event err: %s ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}
