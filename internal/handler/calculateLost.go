package handler

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func CalculateLostBond(ctx *svc.ServiceContext) {
	time.Sleep(30 * time.Second)
	for {
		var disputeGames []schema.DisputeGame
		err := ctx.DB.Where("status!=? and computed=? and calculate_lost=?", schema.DisputeGameStatusInProgress, true, false).Order("block_number").Limit(20).Find(&disputeGames).Error
		if err != nil {
			log.Errorf("[Handler.CalculateLostBond] find dispute games for calculate lost err: %s", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		for _, disputeGame := range disputeGames {
			var credits []schema.GameCredit
			var claimDatas []schema.GameClaimData
			err = ctx.DB.Where("game_contract=?", disputeGame.GameContract).Order("created_at").Find(&credits).Error
			if err != nil {
				log.Errorf("[Handler.CalculateLostBond] find game credit err: %s", errors.WithStack(err))
				panic(err)
			}
			err = ctx.DB.Where("game_contract=?", disputeGame.GameContract).Order("data_index").Find(&claimDatas).Error
			if err != nil {
				log.Errorf("[Handler.CalculateLostBond] find game claim datas err: %s", errors.WithStack(err))
				panic(err)
			}
			for _, credit := range credits {
				address := credit.Address
				amount := credit.Credit
				am := cast.ToInt64(amount)
				for _, claimData := range claimDatas {
					if address == claimData.Claimant {
						am -= cast.ToInt64(claimData.Bond)
					}
				}
				if am < 0 {
					calculateLost := &schema.GameLostBond{
						GameContract: disputeGame.GameContract,
						Address:      address,
						Bond:         cast.ToString(am * -1),
					}
					ctx.DB.Save(calculateLost)
				}
			}
			disputeGame.CalculateLost = true
			ctx.DB.Save(disputeGame)
		}
		time.Sleep(5 * time.Second)
	}
}
