package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cast"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"golang.org/x/time/rate"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/pkg/contract"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"gorm.io/gorm"
)

type RetryDisputeGameClient struct {
	Client             *contract.RateAndRetryDisputeGameClient
	DB                 *gorm.DB
	DisputeGameAddress common.Address
}

func NewRetryDisputeGameClient(db *gorm.DB, address common.Address, rpc *ethclient.Client, limit rate.Limit,
	burst int,
) (*RetryDisputeGameClient, error) {
	newDisputeGame, err := contract.NewDisputeGame(address, rpc)
	if err != nil {
		return nil, err
	}
	retryLimitGame := contract.NewRateAndRetryDisputeGameClient(newDisputeGame, limit, burst)
	return &RetryDisputeGameClient{Client: retryLimitGame, DB: db, DisputeGameAddress: address}, nil
}

func (r *RetryDisputeGameClient) ProcessDisputeGameCreated(ctx context.Context, evt schema.SyncEvent) error {
	dispute := event.DisputeGameCreated{}
	err := dispute.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[processDisputeGameCreated] event data to DisputeGameCreated err: %s", err)
	}
	err = r.addDisputeGame(ctx, &evt)
	if err != nil {
		return fmt.Errorf("[processDisputeGameCreated] addDisputeGame err: %s", err)
	}
	return nil
}

func (r *RetryDisputeGameClient) ProcessDisputeGameMove(ctx context.Context, evt schema.SyncEvent) error {
	disputeGameMove := event.DisputeGameMove{}
	err := disputeGameMove.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] event data to disputeGameMove err: %s", err)
	}
	var storageClaimSize int64
	r.DB.Model(&schema.GameClaimData{}).Where("game_contract=? and on_chain_status = ?",
		evt.ContractAddress, schema.GameClaimDataOnChainStatusValid).Count(&storageClaimSize)
	data, err := r.Client.RetryClaimData(ctx, &bind.CallOpts{}, big.NewInt(storageClaimSize))
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] contract: %s, index: %d move event get claim data err: %s",
			evt.ContractAddress, storageClaimSize, errors.WithStack(err))
	}

	pos := types.NewPositionFromGIndex(data.Position)
	splitDepth, err := r.Client.RetrySplitDepth(ctx, &bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] contract: %s, get splitDepth error: %s", evt.ContractAddress, errors.WithStack(err))
	}
	splitDepths := types.Depth(splitDepth.Uint64())

	poststateBlock, err := r.Client.RetryL2BlockNumber(ctx, &bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] GET game poststateBlock err: %s", err)
	}

	prestateBlock, err := r.Client.RetryStartingBlockNumber(ctx, &bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[processDisputeGameMove] GET prestateBlock err: %s", err)
	}
	var outputblock uint64
	if pos.Depth() > splitDepths {
		outputblock = 0
	} else {
		outputblock, _ = claimedBlockNumber(pos, splitDepths, prestateBlock.Uint64(), poststateBlock.Uint64())
	}

	claimData := &schema.GameClaimData{
		GameContract: evt.ContractAddress,
		DataIndex:    storageClaimSize,
		ParentIndex:  data.ParentIndex,
		CounteredBy:  data.CounteredBy.Hex(),
		Claimant:     data.Claimant.Hex(),
		Bond:         cast.ToString(data.Bond),
		Claim:        hex.EncodeToString(data.Claim[:]),
		Position:     cast.ToString(data.Position),
		Clock:        cast.ToString(data.Clock),
		OutputBlock:  outputblock,
		EventID:      evt.ID,
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

func (r *RetryDisputeGameClient) ProcessDisputeGameResolve(evt schema.SyncEvent) error {
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
	claimData, err := r.Client.RetryClaimData(ctx, &bind.CallOpts{}, big.NewInt(0))
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET index 0 ClaimData err: %s", err)
	}

	gameClaim := &schema.GameClaimData{
		GameContract:  strings.ToLower(disputeGame.DisputeProxy),
		DataIndex:     0,
		ParentIndex:   claimData.ParentIndex,
		CounteredBy:   claimData.CounteredBy.Hex(),
		Claimant:      claimData.Claimant.Hex(),
		Bond:          cast.ToString(claimData.Bond),
		Claim:         hex.EncodeToString(claimData.Claim[:]),
		Position:      cast.ToString(claimData.Position),
		Clock:         cast.ToString(claimData.Clock),
		OutputBlock:   l2Block.Uint64(),
		EventID:       evt.ID,
		OnChainStatus: schema.GameClaimDataOnChainStatusValid,
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
		Status:          schema.DisputeGameStatusInProgress,
		OnChainStatus:   schema.DisputeGameOnChainStatusValid,
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

func (r *RetryDisputeGameClient) ProcessDisputeGameCredit(ctx context.Context) error {
	addresses, err := r.GetAllAddress(&r.DisputeGameAddress)
	if err != nil {
		return fmt.Errorf("[ProcessDisputeGameCredit] GetAllAddress err: %s", err)
	}
	credits := make([]*schema.GameCredit, 0)
	for key := range addresses {
		res, err := r.Client.RetryCredit(ctx, &bind.CallOpts{}, common.HexToAddress(key))
		if err != nil {
			return fmt.Errorf("[handler.SyncCredit] RetryCredit err: %s", err)
		}
		credit := &schema.GameCredit{
			GameContract: r.DisputeGameAddress.String(),
			Address:      key,
			Credit:       res.String(),
		}
		credits = append(credits, credit)
	}
	disputeGame := &schema.DisputeGame{}
	err = r.DB.Where("game_contract=?", strings.ToLower(r.DisputeGameAddress.String())).First(&disputeGame).Error
	if err != nil {
		return fmt.Errorf("[ProcessDisputeGameCredit] find dispute game err:%s", errors.WithStack(err))
	}
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(credits).Error
		if err != nil {
			return fmt.Errorf("[ProcessDisputeGameCredit] save game credit err: %s\n ", err)
		}
		disputeGame.Computed = true
		err = tx.Save(disputeGame).Error
		if err != nil {
			return fmt.Errorf("[ProcessDisputeGameCredit] update dispute game computed err: %s\n ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RetryDisputeGameClient) GetCredit(ctx context.Context, address common.Address) (*big.Int, error) {
	credit, err := r.Client.RetryCredit(ctx, &bind.CallOpts{}, address)
	if err != nil {
		return nil, fmt.Errorf("[GetCredit] GET address %s credit err: %s", address.String(), err)
	}
	return credit, nil
}

func (r *RetryDisputeGameClient) GetAllAddress(disputeGame *common.Address) (map[string]bool, error) {
	claimDatas := make([]schema.GameClaimData, 0)
	err := r.DB.Where("game_contract=?", strings.ToLower(disputeGame.String())).Order("data_index").Find(&claimDatas).Error
	if err != nil {
		return nil, fmt.Errorf("[GetAllAddress] %s find claim datas err: %s", disputeGame.String(), err)
	}
	addresses := map[string]bool{}
	for _, claimData := range claimDatas {
		if claimData.CounteredBy != common.HexToAddress("0").String() {
			if !addresses[claimData.CounteredBy] {
				addresses[claimData.CounteredBy] = true
			}
		}
		if claimData.Claimant != common.HexToAddress("0").String() {
			if !addresses[claimData.Claimant] {
				addresses[claimData.Claimant] = true
			}
		}
	}
	return addresses, nil
}

func claimedBlockNumber(pos types.Position, gameDepth types.Depth, prestateBlock, poststateBlock uint64) (uint64, error) {
	traceIndex := pos.TraceIndex(gameDepth)
	if !traceIndex.IsUint64() {
		return 0, fmt.Errorf("trace index is greater than max uint64: %v", traceIndex)
	}
	outputBlock := traceIndex.Uint64() + prestateBlock + 1
	if outputBlock > poststateBlock {
		outputBlock = poststateBlock
	}
	return outputBlock, nil
}
