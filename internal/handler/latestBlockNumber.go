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
		// use unified RPC manager to get latest block number (automatically applies rate limiting)
		latest, err := ctx.RPCManager.GetLatestBlockNumber(context.Background(), true) // true indicates L1
		if err != nil {
			log.Errorf("[Handler.LatestBlackNumber] Get latest block number error (with rate limit): %s\n", errors.WithStack(err))
			time.Sleep(12 * time.Second)
			continue
		}

		ctx.LatestBlockNumber = cast.ToInt64(latest)
		log.Infof("[Handler.LatestBlackNumber] Latest block number: %d (via RPC Manager)\n", latest)
		time.Sleep(12 * time.Second)
	}
}
