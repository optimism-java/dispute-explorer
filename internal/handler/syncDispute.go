package handler

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	evt "github.com/optimism-java/dispute-explorer/pkg/event"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"golang.org/x/time/rate"
)

func SyncDispute(ctx *svc.ServiceContext) {
	for {
		var events []schema.SyncEvent
		err := ctx.DB.Where("status=? OR status=?", schema.EventPending, schema.EventRollback).Order("block_number").Limit(50).Find(&events).Error
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		if len(events) == 0 {
			log.Infof("[Handler.SyncDispute] Pending events count is 0\n")
			time.Sleep(2 * time.Second)
			continue
		}

		var wg sync.WaitGroup
		for _, event := range events {
			wg.Add(1)
			go func(_wg *sync.WaitGroup, ctx *svc.ServiceContext, event schema.SyncEvent) {
				defer _wg.Done()
				if event.Status == schema.EventPending {
					// add events & block.status= valid
					err = HandlePendingEvent(ctx, event)
					if err != nil {
						log.Errorf("[Handler.SyncEvent] HandlePendingBlock err: %s\n", errors.WithStack(err))
					}
				} else if event.Status == schema.EventRollback {
					// event.status=rollback & block.status=invalid
					err = HandleRollbackEvent(ctx, event)
					if err != nil {
						log.Errorf("[Handler.SyncEvent] HandleRollbackBlock err: %s\n", errors.WithStack(err))
					}
				}
			}(&wg, ctx, event)
		}
		wg.Wait()
	}
}

func HandleRollbackEvent(ctx *svc.ServiceContext, event schema.SyncEvent) error {
	disputeCreated := evt.DisputeGameCreated{}
	disputeMove := evt.DisputeGameMove{}
	disputeResolved := evt.DisputeGameResolved{}
	switch {
	case event.EventName == disputeCreated.Name() && event.EventHash == disputeCreated.EventHash().String():
		// rollback created event include: dispute_game, game_data_claim
		err := disputeCreated.ToObj(event.Data)
		if err != nil {
			log.Errorf("[handle.SyncDispute.RollbackEvent] event data to DisputeGameCreated err: %s", err)
			return errors.WithStack(err)
		}
		// rollback dispute_game
		var disputeGame schema.DisputeGame
		err = ctx.DB.Where("game_contract=?", strings.ToLower(disputeCreated.DisputeProxy)).First(&disputeGame).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("[handle.SyncDispute.RollbackEvent] rollback created event err: %s", err)
		}
		disputeGame.OnChainStatus = schema.DisputeGameOnChainStatusRollBack

		// rollback game_claim_data
		var gameDataClaim schema.GameClaimData
		err = ctx.DB.Where("game_contract=? and data_index=0", strings.ToLower(disputeCreated.DisputeProxy)).First(&gameDataClaim).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("[handle.SyncDispute.RollbackEvent] rollback created the first claim data err: %s", err)
		}
		gameDataClaim.OnChainStatus = schema.GameClaimDataOnChainStatusRollBack

		err = ctx.DB.Transaction(func(tx *gorm.DB) error {
			err = tx.Save(disputeGame).Error
			if err != nil {
				return fmt.Errorf("[handle.SyncDispute.RollbackEvent] update dispute game status err: %s\n ", err)
			}
			err = tx.Save(gameDataClaim).Error
			if err != nil {
				return fmt.Errorf("[handle.SyncDispute.RollbackEvent] update game data claim status err: %s\n ", err)
			}

			event.Status = schema.EventValid
			err = tx.Save(event).Error
			if err != nil {
				return fmt.Errorf("[handle.SyncDispute.RollbackEvent] update event err: %s\n ", err)
			}
			return nil
		})
		// remove contract
		blockchain.RemoveContract(event.ContractAddress)
		log.Infof("remove contract: %s", event.ContractAddress)

	case event.EventName == disputeMove.Name() && event.EventHash == disputeMove.EventHash().String():
		// rollback move: rollback move depend on event_id
		now := time.Now()
		err := ctx.DB.Model(schema.GameClaimData{}).Where("event_id=?", event.ID).
			Updates(map[string]interface{}{"on_chain_status": schema.GameClaimDataOnChainStatusRollBack, "updated_at": now}).Error
		if err != nil {
			log.Errorf("[Handler.SyncDispute.RollbackBlock] rollback move event err: %s ,id : %d \n", err, event.ID)
		}
	case event.EventName == disputeResolved.Name() && event.EventHash == disputeResolved.EventHash().String():
		// rollback resolved
		now := time.Now()
		err := ctx.DB.Model(schema.DisputeGame{}).Where("game_contract=?", event.ContractAddress).
			Updates(map[string]interface{}{"status": schema.DisputeGameStatusInProgress, "updated_at": now}).Error
		if err != nil {
			log.Errorf("[Handler.SyncDispute.RollbackBlock] rollback resolved event err: %s ,id : %d \n", err, event.ID)
		}
		blockchain.AddContract(event.ContractAddress)
		log.Infof("[Handler.SyncDispute.RollbackBlock] rollback resolved event id : %d, contract: %s", event.ID, event.ContractAddress)
	default:
		log.Infof("[Handler.SyncDispute.RollbackBlock] this event does not be monitored %s, hash %s", event.EventName, event.EventHash)
		return nil
	}
	return nil
}

func HandlePendingEvent(ctx *svc.ServiceContext, event schema.SyncEvent) error {
	disputeCreated := evt.DisputeGameCreated{}
	disputeMove := evt.DisputeGameMove{}
	disputeResolved := evt.DisputeGameResolved{}
	switch {
	case event.EventName == disputeCreated.Name() && event.EventHash == disputeCreated.EventHash().String():
		err := disputeCreated.ToObj(event.Data)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] event data to DisputeGameCreated err: %s", err)
			return errors.WithStack(err)
		}
		disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(disputeCreated.DisputeProxy),
			ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] init client for created err: %s", err)
			return errors.WithStack(err)
		}
		err = disputeClient.ProcessDisputeGameCreated(ctx.Context, event)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] ProcessDisputeGameCreated err: %s", err)
			return errors.WithStack(err)
		}
	case event.EventName == disputeMove.Name() && event.EventHash == disputeMove.EventHash().String():
		disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(event.ContractAddress),
			ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] init client for move err: %s", err)
			return errors.WithStack(err)
		}
		err = disputeClient.ProcessDisputeGameMove(ctx.Context, event)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] ProcessDisputeGameCreated err: %s", err)
			return errors.WithStack(err)
		}
	case event.EventName == disputeResolved.Name() && event.EventHash == disputeResolved.EventHash().String():
		disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(event.ContractAddress),
			ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] init client for resolved err: %s", err)
			return errors.WithStack(err)
		}
		err = disputeClient.ProcessDisputeGameResolve(event)
		if err != nil {
			log.Errorf("[handle.SyncDispute.HandlePendingEvent] ProcessDisputeGameCreated err: %s", err)
			return errors.WithStack(err)
		}
	default:
		log.Infof("[handle.SyncDispute.HandlePendingEvent] this event does not be monitored %s, hash %s", event.EventName, event.EventHash)
		return nil
	}
	return nil
}
