package handler

import (
	"context"
	"time"

	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func LatestBlackNumber(ctx *svc.ServiceContext) {
	for {
		latest, err := ctx.L1RPC.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("[Handler.LatestBlackNumber] Syncing block by number error: %s\n", errors.WithStack(err))
			time.Sleep(12 * time.Second)
			continue
		}

		ctx.LatestBlockNumber = cast.ToInt64(latest)
		log.Infof("[Handle.LatestBlackNumber] Syncing latest block number: %d \n", latest)
		time.Sleep(12 * time.Second)
	}
}
