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

func (h DisputeGameHandler) GetCreditRank(c *gin.Context) {
	limit := cast.ToInt(c.Query("limit"))
	if limit == 0 || limit > 100 {
		limit = 10
	}
	var res []CreditRank
	h.DB.Table("game_credit").Select("sum(credit) amount, address").Group("address").Order("amount desc").Limit(limit).Scan(&res)
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

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
