package handler

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/log"
)

func SyncFrontendMoveTransactions(ctx *svc.ServiceContext) {
	log.Infof("[Handler.SyncFrontendMoveTransactions] Starting sync frontend move transactions...")

	for {
		select {
		case <-ctx.Context.Done():
			log.Infof("[Handler.SyncFrontendMoveTransactions] Context canceled, stopping...")
			return
		default:
			err := processFrontendMoveTransactions(ctx)
			if err != nil {
				log.Errorf("[Handler.SyncFrontendMoveTransactions] Process error: %s", err)
			}
			time.Sleep(30 * time.Second)
		}
	}
}

// processFrontendMoveTransactions
func processFrontendMoveTransactions(ctx *svc.ServiceContext) error {
	var unsyncedTransactions []schema.FrontendMoveTransaction
	err := ctx.DB.Where("is_synced = ?", false).Order("id").Limit(10).Find(&unsyncedTransactions).Error
	if err != nil {
		return err
	}

	if len(unsyncedTransactions) == 0 {
		log.Debugf("[Handler.SyncFrontendMoveTransactions] No unsynced transactions found")
		return nil
	}

	log.Infof("[Handler.SyncFrontendMoveTransactions] Found %d unsynced transactions", len(unsyncedTransactions))

	for i := range unsyncedTransactions {
		transaction := &unsyncedTransactions[i]
		err := syncSingleTransaction(ctx, transaction)
		if err != nil {
			log.Errorf("[Handler.SyncFrontendMoveTransactions] Failed to sync transaction %s: %s", transaction.TxHash, err)
			continue
		}
	}

	return nil
}

// syncSingleTransaction
func syncSingleTransaction(ctx *svc.ServiceContext, transaction *schema.FrontendMoveTransaction) error {
	tx := ctx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. update dispute_game has_frontend_move column to true
	err := tx.Model(&schema.DisputeGame{}).
		Where("game_contract = ?", transaction.GameContract).
		Update("has_frontend_move", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. update game_claim_data is_from_frontend column to true
	err = tx.Model(&schema.GameClaimData{}).
		Where("game_contract = ? AND parent_index = ?", transaction.GameContract, transaction.ParentIndex).
		Update("is_from_frontend", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 3. update frontend_move_transactions is_synced column to true
	err = tx.Model(transaction).Update("is_synced", true).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// tx commit
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	log.Infof("[Handler.SyncFrontendMoveTransactions] Successfully synced transaction %s (game: %s, parent_index: %s)",
		transaction.TxHash, transaction.GameContract, transaction.ParentIndex)

	return nil
}
