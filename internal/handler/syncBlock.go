package handler

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/optimism-java/dispute-explorer/pkg/rpc"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func SyncBlock(ctx *svc.ServiceContext) {
	// 防止服务启停切换时同时存在2个服务同步数据
	time.Sleep(10 * time.Second)
	var syncedBlock schema.SyncBlock
	err := ctx.DB.Where("status = ? or status = ? ", schema.BlockValid, schema.BlockPending).Order("block_number desc").First(&syncedBlock).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	switch err {
	case gorm.ErrRecordNotFound:
		ctx.SyncedBlockNumber = ctx.Config.FromBlockNumber
		ctx.SyncedBlockHash = common.HexToHash(ctx.Config.FromBlockHash)
	default:
		ctx.SyncedBlockNumber = syncedBlock.BlockNumber
		ctx.SyncedBlockHash = common.HexToHash(syncedBlock.BlockHash)
	}

	log.Infof("[Handler.SyncBlock]SyncedBlockNumber: %d", ctx.SyncedBlockNumber)
	log.Infof("[Handler.SyncBlock]SyncedBlockHash:%s", ctx.SyncedBlockHash.String())

	for {
		syncingBlockNumber := ctx.SyncedBlockNumber + 1
		log.Infof("[Handler.SyncBlock] Try to sync block number: %d\n", syncingBlockNumber)

		if syncingBlockNumber > ctx.LatestBlockNumber {
			time.Sleep(3 * time.Second)
			continue
		}

		// block, err := ctx.RPC.BlockByNumber(context.Background(), big.NewInt(syncingBlockNumber))
		blockJSON, err := rpc.HTTPPostJSON("", ctx.Config.L1RPCUrl, "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\""+fmt.Sprintf("0x%X", syncingBlockNumber)+"\", true],\"id\":1}")
		if err != nil {
			log.Errorf("[Handler.SyncBlock] Syncing block by number error: %s\n", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		block := rpc.ParseJSONBlock(string(blockJSON))
		log.Infof("[Handler.SyncBlock] Syncing block number: %d, hash: %v, parent hash: %v \n", block.Number(), block.Hash(), block.ParentHash())

		if common.HexToHash(block.ParentHash()) != ctx.SyncedBlockHash {
			log.Errorf("[Handler.SyncBlock] ParentHash of the block being synchronized is inconsistent: %s \n", ctx.SyncedBlockHash)
			continue
		}

		/* Create SyncBlock start */
		err = ctx.DB.Create(&schema.SyncBlock{
			Miner:       block.Result.Miner,
			Blockchain:  ctx.Config.Blockchain,
			BlockTime:   block.Timestamp(),
			BlockNumber: block.Number(),
			BlockHash:   block.Hash(),
			TxCount:     int64(len(block.Result.Transactions)),
			EventCount:  0,
			ParentHash:  block.ParentHash(),
			Status:      schema.BlockPending,
		}).Error
		if err != nil {
			log.Errorf("[Handler.SyncBlock] DB Create SyncBlock error: %s\n", errors.WithStack(err))
			time.Sleep(1 * time.Second)
			continue
		}
		/* Create SyncBlock end */

		ctx.SyncedBlockNumber = block.Number()
		ctx.SyncedBlockHash = common.HexToHash(block.Hash())
	}
}
