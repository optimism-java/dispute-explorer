package handler

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"golang.org/x/time/rate"
)

func SyncDispute(ctx *svc.ServiceContext) {
	for {
		var events []schema.SyncEvent
		err := ctx.DB.Where("status=?", schema.EventPending).Limit(20).Find(&events).Error
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		for _, evt := range events {
			disputeCreated := event.DisputeGameCreated{}
			disputeMove := event.DisputeGameMove{}
			disputeResolved := event.DisputeGameResolved{}
			switch {
			case evt.EventName == disputeCreated.Name() && evt.EventHash == disputeCreated.EventHash().String():
				err = disputeCreated.ToObj(evt.Data)
				if err != nil {
					log.Errorf("[handle.SyncDispute] event data to DisputeGameCreated err: %s", err)
				}
				disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(disputeCreated.DisputeProxy),
					ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
				if err != nil {
					log.Errorf("[handle.SyncDispute] init client for created err: %s", err)
				}
				err = disputeClient.ProcessDisputeGameCreated(ctx.Context, evt)
				if err != nil {
					log.Errorf("[handle.SyncDispute] ProcessDisputeGameCreated err: %s", err)
				}
			case evt.EventName == disputeMove.Name() && evt.EventHash == disputeMove.EventHash().String():
				disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(evt.ContractAddress),
					ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
				if err != nil {
					log.Errorf("[handle.SyncDispute] init client for move err: %s", err)
				}
				err = disputeClient.ProcessDisputeGameMove(ctx.Context, evt)
				if err != nil {
					log.Errorf("[handle.SyncDispute] ProcessDisputeGameCreated err: %s", err)
				}
			case evt.EventName == disputeResolved.Name() && evt.EventHash == disputeResolved.EventHash().String():
				disputeClient, err := NewRetryDisputeGameClient(ctx.DB, common.HexToAddress(evt.ContractAddress),
					ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
				if err != nil {
					log.Errorf("[handle.SyncDispute] init client for resolved err: %s", err)
				}
				err = disputeClient.ProcessDisputeGameResolve(evt)
				if err != nil {
					log.Errorf("[handle.SyncDispute] ProcessDisputeGameCreated err: %s", err)
				}
			default:
				log.Infof("this event does not be monitored %s, hash %s", evt.EventName, evt.EventHash)
			}
		}
	}
}
