package handler

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
)

// SyncClaimDataLen Compensation processing historical data
func SyncClaimDataLen(ctx *svc.ServiceContext) {
	for {
		var disputeGames []schema.DisputeGame
		err := ctx.DB.Where("get_len_status = ? and status != ?",
			false, schema.DisputeGameStatusInProgress).Order("block_number").Limit(50).Find(&disputeGames).Error
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		if len(disputeGames) == 0 {
			log.Debugf("[Handler.SyncClaimDataLen] Pending games count is 0\n")
			time.Sleep(2 * time.Second)
			continue
		}

		for _, disputeGame := range disputeGames {
			var claimDataLen int64
			ctx.DB.Model(schema.GameClaimData{}).Where("game_contract = ? and on_chain_status = ?",
				disputeGame.GameContract, schema.GameClaimDataOnChainStatusValid).Count(&claimDataLen)

			disputeGame.ClaimDataLen = claimDataLen
			disputeGame.GetLenStatus = true

			err := ctx.DB.Save(disputeGame).Error
			if err != nil {
				log.Errorf("[Handler.SyncClaimDataLen] update claim len err", errors.WithStack(err))
			}
		}
		time.Sleep(3 * time.Second)
	}
}
