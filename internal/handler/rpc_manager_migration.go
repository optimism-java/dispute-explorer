package handler

import (
	"context"
	"math/big"
	"time"

	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/optimism-java/dispute-explorer/pkg/rpc"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// LatestBlockNumberWithRateLimit uses unified RPC manager for latest block number retrieval
func LatestBlockNumberWithRateLimit(ctx *svc.ServiceContext) {
	for {
		// use unified RPC manager to get L1 latest block number (with rate limiting)
		latest, err := ctx.RPCManager.GetLatestBlockNumber(context.Background(), true)
		if err != nil {
			log.Errorf("[Handler.LatestBlockNumberWithRateLimit] Get latest block number error: %s\n", errors.WithStack(err))
			time.Sleep(12 * time.Second)
			continue
		}

		ctx.LatestBlockNumber = cast.ToInt64(latest)
		log.Infof("[Handler.LatestBlockNumberWithRateLimit] Latest block number: %d (using RPC Manager)\n", latest)
		time.Sleep(12 * time.Second)
	}
}

// SyncBlockWithRateLimit uses unified RPC manager for block synchronization
func SyncBlockWithRateLimit(ctx *svc.ServiceContext) {
	for {
		// Check pending block count
		// ... existing check logic ...

		syncingBlockNumber := ctx.SyncedBlockNumber + 1
		log.Infof("[Handler.SyncBlockWithRateLimit] Try to sync block number: %d using RPC Manager\n", syncingBlockNumber)

		if syncingBlockNumber > ctx.LatestBlockNumber {
			time.Sleep(3 * time.Second)
			continue
		}

		// Use unified RPC manager to get block (automatically handles rate limiting)
		block, err := ctx.RPCManager.GetBlockByNumber(context.Background(), big.NewInt(syncingBlockNumber), true)
		if err != nil {
			log.Errorf("[Handler.SyncBlockWithRateLimit] Get block by number error: %s\n", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}

		log.Infof("[Handler.SyncBlockWithRateLimit] Got block number: %d, hash: %v, parent hash: %v\n",
			block.Number().Int64(), block.Hash().Hex(), block.ParentHash().Hex())

		// Verify parent hash
		if block.ParentHash() != ctx.SyncedBlockHash {
			log.Errorf("[Handler.SyncBlockWithRateLimit] ParentHash mismatch: expected %s, got %s\n",
				ctx.SyncedBlockHash.Hex(), block.ParentHash().Hex())
			// Can call rollback logic here
			time.Sleep(3 * time.Second)
			continue
		}

		// Save block to database
		// ... existing database save logic ...

		// Update sync status
		ctx.SyncedBlockNumber = block.Number().Int64()
		ctx.SyncedBlockHash = block.Hash()

		log.Infof("[Handler.SyncBlockWithRateLimit] Successfully synced block %d\n", block.Number().Int64())
	}
}

// GetBlockByNumberHTTP gets block using HTTP method (with rate limiting)
func GetBlockByNumberHTTP(ctx *svc.ServiceContext, blockNumber int64) ([]byte, error) {
	requestBody := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"" +
		cast.ToString(blockNumber) + "\", true],\"id\":1}"

	// Use unified RPC manager for HTTP calls (automatically handles rate limiting)
	return ctx.RPCManager.HTTPPostJSON(context.Background(), requestBody, true)
}

// MigrateExistingHandlers example of migrating existing handlers
func MigrateExistingHandlers(ctx *svc.ServiceContext) {
	log.Infof("[Handler.Migration] Starting migration to RPC Manager")

	// Start handlers with rate limiting
	go LatestBlockNumberWithRateLimit(ctx)
	go SyncBlockWithRateLimit(ctx)

	// Start RPC monitoring
	go StartRPCMonitoring(ctx)

	log.Infof("[Handler.Migration] All handlers migrated to use RPC Manager")
}

// StartRPCMonitoring starts RPC monitoring
func StartRPCMonitoring(ctx *svc.ServiceContext) {
	// Create monitor
	monitor := rpc.NewMonitor(ctx.RPCManager, 30*time.Second)

	// Start monitoring
	monitor.Start(ctx.Context)
}

// Compatibility functions: provide smooth migration for existing code
func GetL1Client(ctx *svc.ServiceContext) interface{} {
	// Return rate-limited client wrapper
	return ctx.RPCManager.GetRawClient(true)
}

func GetL2Client(ctx *svc.ServiceContext) interface{} {
	// Return rate-limited client wrapper
	return ctx.RPCManager.GetRawClient(false)
}
