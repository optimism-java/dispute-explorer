package v8

import (
	"fmt"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"gorm.io/gorm"
)

var AddFrontendMoveTrackingTable = gormigrate.Migration{
	ID:      "20240805_add_frontend_move_tracking",
	Migrate: AddFrontendMoveTracking,
}

func AddFrontendMoveTracking(db *gorm.DB) error {
	// Create new frontend_move_transactions table
	err := db.AutoMigrate(&schema.FrontendMoveTransaction{})
	if err != nil {
		return fmt.Errorf("failed to create frontend_move_transactions table: %v", err)
	}

	// Check and add has_frontend_move field to sync_blocks table
	if !db.Migrator().HasColumn(&schema.DisputeGame{}, "has_frontend_move") {
		err = db.Migrator().AddColumn(&schema.DisputeGame{}, "has_frontend_move")
		if err != nil {
			return fmt.Errorf("failed to add has_frontend_move column to dispute_game: %v", err)
		}

		// Set default value for newly added field
		err = db.Exec("UPDATE dispute_game SET has_frontend_move = FALSE WHERE has_frontend_move IS NULL").Error
		if err != nil {
			return fmt.Errorf("failed to set default value for has_frontend_move: %v", err)
		}

		// Create index for new field
		err = db.Exec("CREATE INDEX idx_dispute_game_has_frontend_move ON dispute_game(has_frontend_move)").Error
		if err != nil {
			return fmt.Errorf("failed to create index on has_frontend_move: %v", err)
		}
	}

	// Check and add is_from_frontend field to game_claim_data table
	if !db.Migrator().HasColumn(&schema.GameClaimData{}, "is_from_frontend") {
		err = db.Migrator().AddColumn(&schema.GameClaimData{}, "is_from_frontend")
		if err != nil {
			return fmt.Errorf("failed to add is_from_frontend column to game_claim_data: %v", err)
		}

		// Set default value for newly added field
		err = db.Exec("UPDATE game_claim_data SET is_from_frontend = FALSE WHERE is_from_frontend IS NULL").Error
		if err != nil {
			return fmt.Errorf("failed to set default value for is_from_frontend: %v", err)
		}

		// Create index for new field
		err = db.Exec("CREATE INDEX idx_game_claim_data_is_from_frontend ON game_claim_data(is_from_frontend)").Error
		if err != nil {
			return fmt.Errorf("failed to create index on is_from_frontend: %v", err)
		}
	}

	return nil
}
