package api

import (
	"net/http"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/pkg/util"
	"gorm.io/gorm"
)

type DisputeGameHandler struct {
	DB *gorm.DB
}

func NewDisputeGameHandler(db *gorm.DB) *DisputeGameHandler {
	return &DisputeGameHandler{
		DB: db,
	}
}

// @Summary	Get dispute games
// @schemes
// @Description	Get all dispute game by page
// @Accept			json
// @Produce		json
// @Param			page	query	int	false	"page num"
// @Param			size	query	int	false	"page size"
// @Success		200
// @Router			/disputegames  [get]
func (h DisputeGameHandler) ListDisputeGames(c *gin.Context) {
	games := make([]schema.DisputeGame, 0)
	p := util.NewPagination(c)
	results := h.DB.Order("block_number desc").Find(&games)
	totalRows := results.RowsAffected
	results.Scopes(p.GormPaginate()).Find(&games)
	totalPage := totalRows/p.Size + 1
	c.JSON(http.StatusOK, gin.H{
		"currentPage": p.Page,
		"pageSize":    p.Size,
		"totalCounts": totalRows,
		"totalPage":   totalPage,
		"records":     games,
	})
}

// @Summary	Get claim data
// @schemes
// @Description	Get all claim data by address
// @Accept			json
// @Produce		json
// @Param			address	path	int	false	"dispute game address"
// @Success		200
// @Router			/disputegames/:address/claimdatas  [get]
func (h DisputeGameHandler) GetClaimData(c *gin.Context) {
	address := c.Param("address")
	claimData := make([]schema.GameClaimData, 0)
	h.DB.Where("game_contract = ?", address).Order("data_index").Find(&claimData)
	c.JSON(http.StatusOK, gin.H{
		"data": claimData,
	})
}

type CreditRank struct {
	Amount  string `json:"amount"`
	Address string `json:"address"`
}

// @Summary	Get credit rank
// @schemes
// @Description	Get credit rank
// @Accept			json
// @Produce		json
// @Param			limit	query	int	false	"rank length limit number"
// @Success		200
// @Router			/disputegames/credit/rank  [get]
func (h DisputeGameHandler) GetCreditRank(c *gin.Context) {
	limit := cast.ToInt(c.Query("limit"))
	if limit == 0 || limit > 100 {
		limit = 10
	}
	var res []CreditRank
	h.DB.Table("game_credit").Select("sum(credit) amount, address").Group("address").
		Having("amount!=0").Order("amount desc").Limit(limit).Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// @Summary	Get credit details
// @schemes
// @Description	Get credit details
// @Accept			json
// @Produce		json
// @Param			address	path	int	false	"account address"
// @Success		200
// @Router			/disputegames/:address/credit [get]
func (h DisputeGameHandler) GetCreditDetails(c *gin.Context) {
	address := c.Param("address")
	games := make([]schema.GameCredit, 0)
	p := util.NewPagination(c)
	results := h.DB.Where(" address = ? ", address).Order("created_at desc").Find(&games)
	totalRows := results.RowsAffected
	results.Scopes(p.GormPaginate()).Find(&games)
	totalPage := totalRows/p.Size + 1
	c.JSON(http.StatusOK, gin.H{
		"currentPage": p.Page,
		"pageSize":    p.Size,
		"totalCounts": totalRows,
		"totalPage":   totalPage,
		"records":     games,
	})
}

type Overview struct {
	DisputeGameProxy         string `json:"disputeGameProxy"`
	TotalGames               int64  `json:"totalGames"`
	TotalCredit              string `json:"totalCredit"`
	InProgressGamesCount     int64  `json:"inProgressGamesCount"`
	ChallengerWinGamesCount  int64  `json:"challengerWinGamesCount"`
	DefenderWinWinGamesCount int64  `json:"defenderWinWinGamesCount"`
}

type AmountPerDay struct {
	Amount string `json:"amount"`
	Date   string `json:"date"`
}

// @Summary	overview
// @schemes
// @Description	Get overview
// @Accept			json
// @Produce		json
// @Success		200	{object}	Overview
// @Router			/disputegames/overview [get]
func (h DisputeGameHandler) GetOverview(c *gin.Context) {
	var gameCount int64
	var inProgressGamesCount int64
	var challengerWinGamesCount int64
	var defenderWinWinGamesCount int64
	var totalCredit string
	h.DB.Model(&schema.DisputeGame{}).Count(&gameCount)
	h.DB.Model(&schema.DisputeGame{}).Where("status=?", 0).Count(&inProgressGamesCount)
	h.DB.Model(&schema.DisputeGame{}).Where("status=?", 1).Count(&challengerWinGamesCount)
	h.DB.Model(&schema.DisputeGame{}).Where("status=?", 2).Count(&defenderWinWinGamesCount)
	h.DB.Model(&schema.GameCredit{}).Select("IFNULL(sum(credit), 0) as totalCredit").Find(&totalCredit)
	overview := &Overview{
		DisputeGameProxy:         "0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1",
		TotalGames:               gameCount,
		TotalCredit:              totalCredit,
		InProgressGamesCount:     inProgressGamesCount,
		ChallengerWinGamesCount:  challengerWinGamesCount,
		DefenderWinWinGamesCount: defenderWinWinGamesCount,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": overview,
	})
}

// @Summary	GetAmountPerDays
// @schemes
// @Description	Get amount per day
// @Accept			json
// @Produce		json
// @Param			days	query		int	false	"today before ? days"
// @Success		200		{object}	[]AmountPerDay
// @Router			/disputegames/overview/amountperdays [get]
func (h DisputeGameHandler) GetAmountPerDays(c *gin.Context) {
	days := cast.ToInt(c.Query("days"))
	res := make([]AmountPerDay, 0)
	if days == 0 || days > 100 {
		days = 15
	}
	h.DB.Raw(" select sum(a.bond) amount, DATE_FORMAT(FROM_UNIXTIME(se.block_time),'%Y-%m-%d') date "+
		" from game_claim_data a left join sync_events se on a.event_id = se.id "+
		" where DATE_FORMAT(FROM_UNIXTIME(se.block_time),'%Y-%m-%d') >= DATE_FORMAT((NOW() - INTERVAL ? DAY),'%Y-%m-%d') "+
		" group by date", days).Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// @Summary	GetBondInProgressPerDays
// @schemes
// @Description	Get bond in progress per days
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]AmountPerDay
// @Router			/disputegames/statistics/bond/inprogress [get]
func (h DisputeGameHandler) GetBondInProgressPerDays(c *gin.Context) {
	res := make([]AmountPerDay, 0)
	h.DB.Raw(" select sum(a.bond) amount, DATE_FORMAT(FROM_UNIXTIME(se.block_time),'%Y-%m-%d') date " +
		" from game_claim_data a left join sync_events se on a.event_id = se.id" +
		" left join dispute_game dg on a.game_contract = dg.game_contract where dg.status=0 group by date order by date desc").Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

type CountGames struct {
	Amount string `json:"amount"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

// @Summary	GetCountDisputeGameGroupByStatus
// @schemes
// @Description	Get dispute games count group by status and per day
// @Accept			json
// @Produce		json
// @Param			days	query		int	false	"today before ? days"
// @Success		200		{object}	[]CountGames
// @Router			/disputegames/count [get]
func (h DisputeGameHandler) GetCountDisputeGameGroupByStatus(c *gin.Context) {
	days := cast.ToInt(c.Query("days"))
	res := make([]CountGames, 0)
	if days == 0 || days > 100 {
		days = 15
	}
	h.DB.Raw("select count(1) count, DATE_FORMAT(FROM_UNIXTIME(dg.block_time),'%Y-%m-%d') date, status "+
		" from dispute_game dg where DATE_FORMAT(FROM_UNIXTIME(dg.block_time),'%Y-%m-%d') >= "+
		" DATE_FORMAT((NOW() - INTERVAL ? DAY),'%Y-%m-%d') group by date, status order by date desc", days).Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// @Summary	ListGameEvents
// @schemes
// @Description	Get game events
// @Accept			json
// @Produce		json
// @Param			days	query		int	false	"today before ? days"
// @Success		200
// @Router			/disputegames/events [get]
func (h DisputeGameHandler) ListGameEvents(c *gin.Context) {
	events := make([]schema.SyncEvent, 0)
	p := util.NewPagination(c)
	results := h.DB.Order("block_number desc").Find(&events)
	totalRows := results.RowsAffected
	results.Scopes(p.GormPaginate()).Find(&events)
	totalPage := totalRows/p.Size + 1
	c.JSON(http.StatusOK, gin.H{
		"currentPage": p.Page,
		"pageSize":    p.Size,
		"totalCounts": totalRows,
		"totalPage":   totalPage,
		"records":     events,
	})
}
