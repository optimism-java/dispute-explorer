package handler

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/contract"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"gorm.io/gorm"
)

func BatchFilterAddAndRemove(ctx *svc.ServiceContext, events []*schema.SyncEvent) error {
	for _, evt := range events {
		err := filterAddAndRemove(ctx, evt)
		if err != nil {
			return fmt.Errorf("[BatchFilterAddAndRemove] filterAddAndRemove: %s", err)
		}
	}
	return nil
}

func filterAddAndRemove(ctx *svc.ServiceContext, evt *schema.SyncEvent) error {
	dispute := event.DisputeGameCreated{}
	if evt.EventName == dispute.Name() && evt.EventHash == dispute.EventHash().String() {
		err := dispute.ToObj(evt.Data)
		if err != nil {
			return fmt.Errorf("[FilterDisputeContractAndAdd] event data to DisputeGameCreated err: %s", err)
		}
		err = addDisputeGame(ctx, evt)
		if err != nil {
			return fmt.Errorf("[FilterDisputeContractAndAdd] addDisputeGame err: %s", err)
		}
		blockchain.AddContract(dispute.DisputeProxy)
	}
	disputeResolved := event.DisputeGameResolved{}
	if evt.EventName == disputeResolved.Name() && evt.EventHash == disputeResolved.EventHash().String() {
		err := disputeResolved.ToObj(evt.Data)
		if err != nil {
			return fmt.Errorf("[FilterDisputeContractAndAdd] event data to DisputeGameResolved err: %s", err)
		}
		var game schema.DisputeGame
		err = ctx.DB.Where(" contract_address = ? ", evt.ContractAddress).First(&game).Error
		if err != nil {
			return fmt.Errorf("[FilterDisputeContractAndAdd] resolved event find game err: %s", err)
		}
		game.Status = disputeResolved.Status
		ctx.DB.Save(game)
		blockchain.RemoveContract(evt.ContractAddress)
	}
	disputeGameMove := event.DisputeGameMove{}
	if evt.EventName == disputeGameMove.Name() && evt.EventHash == disputeGameMove.EventHash().String() {
		err := disputeGameMove.ToObj(evt.Data)
		if err != nil {
			return fmt.Errorf("[FilterDisputeContractAndAdd] event data to disputeGameMove err: %s", err)
		}
		newDisputeGame, err := contract.NewDisputeGame(common.HexToAddress(evt.ContractAddress), ctx.L1RPC)
		if err != nil {
			return fmt.Errorf("[addDisputeGame] init dispute game contract client err: %s", err)
		}
		index := disputeGameMove.ParentIndex.Add(disputeGameMove.ParentIndex, big.NewInt(1))
		data, err := newDisputeGame.ClaimData(&bind.CallOpts{}, index)
		if err != nil {
			return fmt.Errorf("[addDisputeGame] contract: %s, index: %d move event get claim data err: %s", evt.ContractAddress, index, err)
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
		ctx.DB.Save(claimData)
	}

	return nil
}

func addDisputeGame(ctx *svc.ServiceContext, evt *schema.SyncEvent) error {
	disputeGame := event.DisputeGameCreated{}
	err := disputeGame.ToObj(evt.Data)
	if err != nil {
		return fmt.Errorf("[addDisputeGame] event data to DisputeGameCreated err: %s", err)
	}
	newDisputeGame, err := contract.NewDisputeGame(common.HexToAddress(disputeGame.DisputeProxy), ctx.L1RPC)
	if err != nil {
		return fmt.Errorf("[addDisputeGame] init dispute game contract client err: %s", err)
	}
	l2Block, err := newDisputeGame.L2BlockNumber(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET game L2BlockNumber err: %s", err)
	}
	status, err := newDisputeGame.Status(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET game status err: %s", err)
	}
	claimData, err := newDisputeGame.ClaimData(&bind.CallOpts{}, big.NewInt(0))
	if err != nil {
		return fmt.Errorf("[addDisputeGame] GET index 0 ClaimData err: %s", err)
	}

	gameClaim := &schema.GameClaimData{
		GameContract: disputeGame.DisputeProxy,
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
		GameContract:    disputeGame.DisputeProxy,
		GameType:        disputeGame.GameType,
		L2BlockNumber:   l2Block.Int64(),
		Status:          status,
	}
	err = ctx.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(game).Error
		if err != nil {
			return fmt.Errorf("[addDisputeGame] save dispute game err: %s\n ", err)
		}
		err = tx.Save(gameClaim).Error
		if err != nil {
			return fmt.Errorf("[addDisputeGame] save game claim err: %s\n ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}
