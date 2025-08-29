package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/optimism-java/dispute-explorer/docs"
	"github.com/optimism-java/dispute-explorer/internal/api"
	"github.com/optimism-java/dispute-explorer/internal/handler"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/internal/types"
	"github.com/optimism-java/dispute-explorer/migration/migrate"
	disputeLog "github.com/optimism-java/dispute-explorer/pkg/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ctx := context.Background()
	cfg := types.GetConfig()
	disputeLog.Init(cfg.LogLevel, cfg.LogFormat)
	disputeLog.Infof("config: %v\n", cfg)
	sCtx := svc.NewServiceContext(ctx, cfg)
	migrate.Migrate(sCtx.DB)
	handler.Run(sCtx)
	disputeLog.Info("listener running...\n")

	router := gin.Default()
	disputeGameHandler := api.NewDisputeGameHandler(sCtx.DB, sCtx.RPCManager, cfg)

	frontendMoveAPI := api.NewFrontendMoveAPI(sCtx)
	docs.SwaggerInfo.Title = "Dispute Game Swagger API"
	docs.SwaggerInfo.Description = "This is a dispute-explorer server."
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/disputegames", disputeGameHandler.ListDisputeGames)
	router.GET("/disputegames/:address/claimdatas", disputeGameHandler.GetClaimData)
	router.GET("/disputegames/credit/rank", disputeGameHandler.GetCreditRank)
	router.GET("/disputegames/:address/credit", disputeGameHandler.GetCreditDetails)
	router.GET("/disputegames/overview", disputeGameHandler.GetOverview)
	router.GET("/disputegames/overview/amountperday", disputeGameHandler.GetAmountPerDays)
	router.GET("/disputegames/statistics/bond/inprogress", disputeGameHandler.GetBondInProgressPerDays)
	router.GET("/disputegames/daylycount", disputeGameHandler.GetCountDisputeGameGroupByStatus)
	router.GET("/disputegames/events", disputeGameHandler.ListGameEvents)
	router.GET("/disputegames/claimroot/:blockNumber", disputeGameHandler.GetClaimRoot)
	router.POST("/disputegames/calculate/claim", disputeGameHandler.GetGamesClaimByPosition)
	router.GET("/disputegames/chainname", disputeGameHandler.GetCurrentBlockChain)

	router.POST("/disputegames/frontend-move", frontendMoveAPI.RecordMove)
	router.GET("/disputegames/:address/frontend-moves", frontendMoveAPI.GetMovesByGame)
	router.GET("/disputegames/frontend-move/:txhash", frontendMoveAPI.GetMoveByTxHash)
	router.GET("/disputegames/with-frontend-flag", frontendMoveAPI.GetGamesWithFrontendFlag)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":" + cfg.APIPort)
	if err != nil {
		disputeLog.Errorf("start error %s", err)
		return
	}
}
