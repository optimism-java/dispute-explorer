package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"

	"gorm.io/gorm"
)

// FrontendMoveRequest request structure for frontend-initiated move transactions
type FrontendMoveRequest struct {
	GameContract   string `json:"game_contract" binding:"required"`   // Dispute Game contract address
	TxHash         string `json:"tx_hash" binding:"required"`         // Transaction hash
	Claimant       string `json:"claimant" binding:"required"`        // Initiator address
	ParentIndex    string `json:"parent_index" binding:"required"`    // Parent index
	Claim          string `json:"claim" binding:"required"`           // Claim data
	IsAttack       bool   `json:"is_attack"`                          // Whether it's an attack
	ChallengeIndex string `json:"challenge_index" binding:"required"` // Challenge index
	DisputedClaim  string `json:"disputed_claim" binding:"required"`  // Disputed claim
}

// FrontendMoveHandler handles frontend-initiated move transactions
type FrontendMoveHandler struct {
	svc *svc.ServiceContext
}

// NewFrontendMoveHandler creates a new FrontendMoveHandler
func NewFrontendMoveHandler(svc *svc.ServiceContext) *FrontendMoveHandler {
	return &FrontendMoveHandler{
		svc: svc,
	}
}

// GetServiceContext gets ServiceContext
func (h *FrontendMoveHandler) GetServiceContext() *svc.ServiceContext {
	return h.svc
}

// RecordFrontendMove records frontend-initiated move transactions
func (h *FrontendMoveHandler) RecordFrontendMove(req *FrontendMoveRequest) error {
	// Validate contract address format
	if !common.IsHexAddress(req.GameContract) {
		return fmt.Errorf("invalid game contract address: %s", req.GameContract)
	}

	// Validate transaction hash format
	if len(req.TxHash) != 66 || !strings.HasPrefix(req.TxHash, "0x") {
		return fmt.Errorf("invalid transaction hash format: %s", req.TxHash)
	}

	// Validate initiator address format
	if !common.IsHexAddress(req.Claimant) {
		return fmt.Errorf("invalid claimant address: %s", req.Claimant)
	}

	// Check if this transaction has already been recorded
	var existingRecord schema.FrontendMoveTransaction
	err := h.svc.DB.Where("tx_hash = ?", req.TxHash).First(&existingRecord).Error
	if err == nil {
		log.Warnf("[FrontendMoveHandler] Transaction %s already recorded", req.TxHash)
		return fmt.Errorf("transaction %s already recorded", req.TxHash)
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing record: %v", err)
	}

	// Create new record
	frontendMove := &schema.FrontendMoveTransaction{
		GameContract:   req.GameContract,
		TxHash:         req.TxHash,
		Claimant:       req.Claimant,
		ParentIndex:    req.ParentIndex,
		Claim:          req.Claim,
		IsAttack:       req.IsAttack,
		ChallengeIndex: req.ChallengeIndex,
		DisputedClaim:  req.DisputedClaim,
		Status:         schema.FrontendMoveStatusPending,
		SubmittedAt:    time.Now().Unix(),
		IsSynced:       false, // 新记录默认未同步
	}

	// Save to database
	err = h.svc.DB.Create(frontendMove).Error
	if err != nil {
		return fmt.Errorf("failed to save frontend move record: %v", err)
	}

	log.Infof("[FrontendMoveHandler] Recorded frontend move transaction: %s for game: %s", req.TxHash, req.GameContract)

	// Asynchronously check transaction status
	go h.monitorTransactionStatus(frontendMove.ID, req.TxHash)

	return nil
}

// monitorTransactionStatus monitors transaction status
func (h *FrontendMoveHandler) monitorTransactionStatus(recordID int64, txHash string) {
	maxRetries := 60 // Maximum 60 retries, 10 seconds interval each
	retryInterval := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		time.Sleep(retryInterval)

		// Query transaction status using RPCManager
		l1Client := h.svc.RPCManager.GetRawClient(true)
		receipt, err := l1Client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Debugf("[FrontendMoveHandler] Transaction %s not yet mined or error: %v", txHash, err)
			continue
		}

		// Update record status
		var status string
		var confirmedAt int64
		if receipt.Status == 1 {
			status = schema.FrontendMoveStatusConfirmed
			confirmedAt = time.Now().Unix()
		} else {
			status = schema.FrontendMoveStatusFailed
		}
		err = h.svc.DB.Model(&schema.FrontendMoveTransaction{}).
			Where("id = ?", recordID).
			Updates(map[string]interface{}{
				"status":       status,
				"block_number": receipt.BlockNumber.Int64(),
				"confirmed_at": confirmedAt,
			}).Error
		if err != nil {
			log.Errorf("[FrontendMoveHandler] Failed to update transaction status for %s: %v", txHash, err)
			return
		}
		log.Infof("[FrontendMoveHandler] Transaction %s status updated to %s", txHash, status)

		// If transaction is successful, mark related records
		// if receipt.Status == 1 {
		//	h.markRelatedRecords(txHash, receipt.BlockNumber.Int64())
		// }

		return
	}

	// If timeout without finding transaction, mark as failed
	err := h.svc.DB.Model(&schema.FrontendMoveTransaction{}).
		Where("id = ?", recordID).
		Updates(map[string]interface{}{
			"status":        schema.FrontendMoveStatusFailed,
			"error_message": "Transaction timeout",
		}).Error
	if err != nil {
		log.Errorf("[FrontendMoveHandler] Failed to update timeout status for %s: %v", txHash, err)
	}

	log.Warnf("[FrontendMoveHandler] Transaction %s monitoring timeout", txHash)
}

// markRelatedRecords marks related block and event records
func (h *FrontendMoveHandler) markRelatedRecords(txHash string, blockNumber int64) {
	// Mark related blocks
	err := h.svc.DB.Model(&schema.DisputeGame{}).
		Where("block_number = ?", blockNumber).
		Update("has_frontend_move", true).Error
	if err != nil {
		log.Errorf("[FrontendMoveHandler] Failed to mark block %d for tx %s: %v", blockNumber, txHash, err)
	}

	// Mark related events
	err = h.svc.DB.Model(&schema.GameClaimData{}).
		Where("tx_hash = ?", txHash).
		Update("is_from_frontend", true).Error
	if err != nil {
		log.Errorf("[FrontendMoveHandler] Failed to mark events for tx %s: %v", txHash, err)
	}

	log.Infof("[FrontendMoveHandler] Marked related records for tx %s in block %d", txHash, blockNumber)
}

// GetFrontendMovesByGame gets frontend-initiated move transactions for specified game
func (h *FrontendMoveHandler) GetFrontendMovesByGame(gameContract string, page, size int) ([]schema.FrontendMoveTransaction, int64, error) {
	var moves []schema.FrontendMoveTransaction
	var total int64

	// Validate contract address format
	if !common.IsHexAddress(gameContract) {
		return nil, 0, fmt.Errorf("invalid game contract address: %s", gameContract)
	}

	// Get total count
	err := h.svc.DB.Model(&schema.FrontendMoveTransaction{}).
		Where("game_contract = ?", gameContract).
		Count(&total).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count frontend moves: %v", err)
	}

	// Get paginated data
	offset := (page - 1) * size
	err = h.svc.DB.Where("game_contract = ?", gameContract).
		Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&moves).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get frontend moves: %v", err)
	}

	return moves, total, nil
}

// GetFrontendMoveByTxHash gets frontend-initiated move transaction by transaction hash
func (h *FrontendMoveHandler) GetFrontendMoveByTxHash(txHash string) (*schema.FrontendMoveTransaction, error) {
	if len(txHash) != 66 || !strings.HasPrefix(txHash, "0x") {
		return nil, fmt.Errorf("invalid transaction hash format: %s", txHash)
	}

	var move schema.FrontendMoveTransaction
	err := h.svc.DB.Where("tx_hash = ?", txHash).First(&move).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("frontend move not found for tx: %s", txHash)
		}
		return nil, fmt.Errorf("failed to get frontend move: %v", err)
	}

	return &move, nil
}
