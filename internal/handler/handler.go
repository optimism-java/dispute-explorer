package handler

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/rpc"
)

func Run(ctx *svc.ServiceContext) {
	// Start RPC monitoring
	go startRPCMonitoring(ctx)

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
	// calculate lost bond
	go CalculateLostBond(ctx)
	// sync claim len
	go SyncClaimDataLen(ctx)
	// sync frontend move transactions
	go SyncFrontendMoveTransactions(ctx)
}

// startRPCMonitoring starts RPC monitoring (internal function)
func startRPCMonitoring(ctx *svc.ServiceContext) {
	// Create monitor, output statistics every 30 seconds
	monitor := rpc.NewMonitor(ctx.RPCManager, 30*time.Second)

	// Start monitoring
	monitor.Start(ctx.Context)
}
