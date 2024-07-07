package handler

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

func SyncCredit(ctx *svc.ServiceContext) {
	time.Sleep(5 * time.Second)
	for {
		var disputeGames []schema.DisputeGame
		err := ctx.DB.Where("status!=? and computed=?", schema.DisputeGameStatusInProgress, false).Order("block_number").Limit(20).Find(&disputeGames).Error
		if err != nil {
			log.Errorf("[Handler.SyncCredit] find dispute games for statistic err: %s", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		for _, disputeGame := range disputeGames {
			game := common.HexToAddress(disputeGame.GameContract)
			disputeClient, err := NewRetryDisputeGameClient(ctx.DB, game,
				ctx.L1RPC, rate.Limit(ctx.Config.RPCRateLimit), ctx.Config.RPCRateBurst)
			if err != nil {
				log.Errorf("[Handler.SyncCredit] NewRetryDisputeGameClient err: %s", err)
				time.Sleep(5 * time.Second)
				continue
			}
			err = disputeClient.ProcessDisputeGameCredit(ctx.Context)
			if err != nil {
				log.Errorf("[Handler.SyncCredit] ProcessDisputeGameCredit err: %s", err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}
