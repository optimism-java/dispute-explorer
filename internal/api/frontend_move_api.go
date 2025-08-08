package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/optimism-java/dispute-explorer/internal/handler"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"gorm.io/gorm"
)

type FrontendMoveAPI struct {
	handler *handler.FrontendMoveHandler
}

// NewFrontendMoveAPI creates a new FrontendMoveAPI
func NewFrontendMoveAPI(svc *svc.ServiceContext) *FrontendMoveAPI {
	return &FrontendMoveAPI{
		handler: handler.NewFrontendMoveHandler(svc),
	}
}

// RecordMoveResponse response for recording Move transactions
type RecordMoveResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// FrontendMovesResponse response for frontend-initiated Move transaction list
type FrontendMovesResponse struct {
	Success bool                             `json:"success"`
	Data    []schema.FrontendMoveTransaction `json:"data"`
	Total   int64                            `json:"total"`
	Page    int                              `json:"page"`
	Size    int                              `json:"size"`
}

// FrontendMoveDetailResponse response for frontend-initiated Move transaction details
type FrontendMoveDetailResponse struct {
	Success bool                            `json:"success"`
	Data    *schema.FrontendMoveTransaction `json:"data,omitempty"`
	Message string                          `json:"message,omitempty"`
}

// @Summary	Record frontend move transaction
// @schemes
// @Description	Record a move transaction initiated from frontend
// @Accept			json
// @Produce		json
// @Param			request	body	handler.FrontendMoveRequest	true	"Frontend move request"
// @Success		200	{object}	RecordMoveResponse
// @Router			/disputegames/frontend-move  [post]
func (api *FrontendMoveAPI) RecordMove(c *gin.Context) {
	var req handler.FrontendMoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("[FrontendMoveAPI] Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, RecordMoveResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
		return
	}

	err := api.handler.RecordFrontendMove(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("[FrontendMoveAPI] Failed to record frontend move: %v", err)
		c.JSON(http.StatusInternalServerError, RecordMoveResponse{
			Success: false,
			Message: "Failed to record move transaction: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, RecordMoveResponse{
		Success: true,
		Message: "Move transaction recorded successfully",
	})
}

// @Summary	Get frontend moves by game
// @schemes
// @Description	Get all frontend move transactions for a specific dispute game
// @Accept			json
// @Produce		json
// @Param			address	path	string	true	"Dispute game contract address"
// @Param			page	query	int	false	"Page number (default: 1)"
// @Param			size	query	int	false	"Page size (default: 10)"
// @Success		200	{object}	FrontendMovesResponse
// @Router			/disputegames/:address/frontend-moves  [get]
func (api *FrontendMoveAPI) GetMovesByGame(c *gin.Context) {
	gameContract := c.Param("address")
	if gameContract == "" {
		c.JSON(http.StatusBadRequest, FrontendMovesResponse{
			Success: false,
		})
		return
	}

	// Parse pagination parameters
	page := 1
	size := 10
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if sizeStr := c.Query("size"); sizeStr != "" {
		if s, err := strconv.Atoi(sizeStr); err == nil && s > 0 && s <= 100 {
			size = s
		}
	}

	moves, total, err := api.handler.GetFrontendMovesByGame(gameContract, page, size)
	if err != nil {
		log.Errorf("[FrontendMoveAPI] Failed to get frontend moves: %v", err)
		c.JSON(http.StatusInternalServerError, FrontendMovesResponse{
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, FrontendMovesResponse{
		Success: true,
		Data:    moves,
		Total:   total,
		Page:    page,
		Size:    size,
	})
}

// @Summary	Get frontend move by transaction hash
// @schemes
// @Description	Get frontend move transaction details by transaction hash
// @Accept			json
// @Produce		json
// @Param			txhash	path	string	true	"Transaction hash"
// @Success		200	{object}	FrontendMoveDetailResponse
// @Router			/disputegames/frontend-move/:txhash  [get]
func (api *FrontendMoveAPI) GetMoveByTxHash(c *gin.Context) {
	txHash := c.Param("txhash")
	if txHash == "" {
		c.JSON(http.StatusBadRequest, FrontendMoveDetailResponse{
			Success: false,
			Message: "Transaction hash is required",
		})
		return
	}

	move, err := api.handler.GetFrontendMoveByTxHash(txHash)
	if err != nil {
		log.Errorf("[FrontendMoveAPI] Failed to get frontend move: %v", err)
		statusCode := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, FrontendMoveDetailResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FrontendMoveDetailResponse{
		Success: true,
		Data:    move,
	})
}

// @Summary	Get dispute games with frontend move flag
// @schemes
// @Description	Get all dispute games with information about whether they contain frontend-initiated moves
// @Accept			json
// @Produce		json
// @Param			page	query	int	false	"Page number (default: 1)"
// @Param			size	query	int	false	"Page size (default: 10)"
// @Param			frontend_only	query	bool	false	"Only show games with frontend moves"
// @Success		200
// @Router			/disputegames/with-frontend-flag  [get]
func (api *FrontendMoveAPI) GetGamesWithFrontendFlag(c *gin.Context) {
	// Parse pagination parameters
	page := 1
	size := 10
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if sizeStr := c.Query("size"); sizeStr != "" {
		if s, err := strconv.Atoi(sizeStr); err == nil && s > 0 && s <= 100 {
			size = s
		}
	}

	frontendOnly := c.Query("frontend_only") == "true"

	// Build query
	query := api.handler.GetServiceContext().DB.Model(&schema.DisputeGame{}).
		Select("dispute_games.*, COALESCE(frontend_stats.has_frontend_move, false) as has_frontend_move").
		Joins(`LEFT JOIN (
			SELECT game_contract, true as has_frontend_move 
			FROM frontend_move_transactions 
			WHERE status = ? 
			GROUP BY game_contract
		) frontend_stats ON dispute_games.game_contract = frontend_stats.game_contract`, schema.FrontendMoveStatusConfirmed)

	if frontendOnly {
		query = query.Where("frontend_stats.has_frontend_move = true")
	}

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.Errorf("[FrontendMoveAPI] Failed to count games: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to count games",
		})
		return
	}

	var games []map[string]interface{}
	offset := (page - 1) * size
	err = query.Offset(offset).Limit(size).Order("created_at DESC").Find(&games).Error
	if err != nil {
		log.Errorf("[FrontendMoveAPI] Failed to get games: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get games",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    games,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}
