package handler

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func LogFilter(ctx *svc.ServiceContext, block schema.SyncBlock, addresses []common.Address, topics [][]common.Hash) ([]*schema.SyncEvent, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block.BlockNumber),
		ToBlock:   big.NewInt(block.BlockNumber),
		Topics:    topics,
		Addresses: addresses,
	}
	logs, err := ctx.L1RPC.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Infof("[CancelOrder.Handle] Cancel Pending List Length is %d ,block number is %d \n", len(logs), block.BlockNumber)
	return LogsToEvents(ctx, logs, block.ID)
}

func LogsToEvents(ctx *svc.ServiceContext, logs []types.Log, syncBlockID int64) ([]*schema.SyncEvent, error) {
	events := []*schema.SyncEvent{}
	blockTimes := make(map[int64]int64)
	for _, vlog := range logs {
		eventHash := event.TopicToHash(vlog, 0)
		contractAddress := vlog.Address
		Event := blockchain.GetEvent(eventHash)
		if Event == nil {
			log.Infof("[LogsToEvents] logs[txHash: %s, contractAddress:%s, eventHash: %s]\n", vlog.TxHash, strings.ToLower(contractAddress.Hex()), eventHash)
			continue
		}

		blockTime := blockTimes[cast.ToInt64(vlog.BlockNumber)]
		if blockTime == 0 {
			block, err := ctx.L1RPC.BlockByNumber(context.Background(), big.NewInt(cast.ToInt64(vlog.BlockNumber)))
			if err != nil {
				return nil, errors.WithStack(err)
			}
			blockTime = cast.ToInt64(block.Time())
		}
		data, err := Event.Data(vlog)
		if err != nil {
			log.Errorf("[LogsToEvents] logs[txHash: %s, contractAddress:%s, eventHash: %s]\n", vlog.TxHash, strings.ToLower(contractAddress.Hex()), eventHash)
			log.Errorf("[LogsToEvents] data err: %s\n", errors.WithStack(err))
			continue
		}

		events = append(events, &schema.SyncEvent{
			Blockchain:      ctx.Config.Blockchain,
			SyncBlockID:     syncBlockID,
			BlockTime:       blockTime,
			BlockNumber:     cast.ToInt64(vlog.BlockNumber),
			BlockHash:       vlog.BlockHash.Hex(),
			BlockLogIndexed: cast.ToInt64(vlog.Index),
			TxIndex:         cast.ToInt64(vlog.TxIndex),
			TxHash:          vlog.TxHash.Hex(),
			EventName:       Event.Name(),
			EventHash:       eventHash.Hex(),
			ContractAddress: strings.ToLower(contractAddress.Hex()),
			Data:            data,
			Status:          schema.EventPending,
		})
	}
	return events, nil
}
