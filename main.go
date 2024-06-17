package main

import (
	"github.com/gin-gonic/gin"
	"github.com/optimism-java/dispute-explorer/internal/api"
	"github.com/optimism-java/dispute-explorer/internal/handler"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/internal/types"
	"github.com/optimism-java/dispute-explorer/pkg/log"
)

func main() {
	cfg := types.GetConfig()
	log.Init(cfg.LogLevel, cfg.LogFormat)
	log.Infof("config: %v\n", cfg)
	ctx := svc.NewServiceContext(cfg)
	handler.Run(ctx)
	log.Info("listener running...\n")
	router := gin.Default()

	disputeGameHandler := api.NewDisputeGameHandler(ctx.DB)

	router.GET("/games", disputeGameHandler.ListDisputeGames)
	router.GET("/games/:address/claim-data", disputeGameHandler.GetClaimData)

	err := router.Run()
	if err != nil {
		log.Errorf("start error %s", err)
		return
	}

}
