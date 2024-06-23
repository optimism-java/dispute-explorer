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
