package handler

import (
	"time"

	"github.com/optimism-java/dispute-explorer/pkg/rpc"
	"github.com/pkg/errors"

	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
)

func LatestBlackNumber(ctx *svc.ServiceContext) {
	for {
		blockJSON, err := rpc.HTTPPostJSON("", ctx.Config.L1RPCUrl, "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"safe\", false],\"id\":1}")
		if err != nil {
			log.Errorf("[Handler.LatestBlackNumber] Syncing block by number error: %s\n", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		block := rpc.ParseJSONBlock(string(blockJSON))

		ctx.LatestBlockNumber = block.Number()
		log.Infof("[Handle.LatestBlackNumber] Syncing latest block number: %d \n", block.Number())
		time.Sleep(3 * time.Second)
	}
}
