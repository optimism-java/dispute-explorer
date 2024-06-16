package handler

import (
	"context"
	"time"

	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
)

func LatestBlackNumber(ctx *svc.ServiceContext) {
	for {
		latest, err := ctx.L1RPC.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("[Handle.LatestBlackNumber]Syncing latest block number error: %s\n", err)
			time.Sleep(3 * time.Second)
			continue
		}
		ctx.LatestBlockNumber = int64(latest)
		log.Infof("[Handle.LatestBlackNumber] Syncing latest block number: %d \n", latest)
		time.Sleep(3 * time.Second)
	}
}
