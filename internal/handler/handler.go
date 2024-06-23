package handler

import (
	"github.com/optimism-java/dispute-explorer/internal/svc"
)

func Run(ctx *svc.ServiceContext) {
	// query last block number
	go LatestBlackNumber(ctx)
	// sync blocks
	go SyncBlock(ctx)
	// sync events
	go SyncEvent(ctx)
	// sync dispute game
	go SyncDispute(ctx)
	// sync credit
	go SyncCredit(ctx)
}
