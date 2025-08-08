package v8

import (
	"fmt"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"gorm.io/gorm"
)

var AddFrontendMoveTrackingTable = gormigrate.Migration{
	ID: "20240805_add_frontend_move_tracking",
	Migrate: func(db *gorm.DB) error {
		return AddFrontendMoveTracking(db)
	},
}

func AddFrontendMoveTracking(db *gorm.DB) error {
	// Create new frontend_move_transactions table
	err := db.AutoMigrate(&schema.FrontendMoveTransaction{})
	if err != nil {
		return fmt.Errorf("failed to create frontend_move_transactions table: %v", err)
	}

	// Check and add has_frontend_move field to sync_blocks table
	if !db.Migrator().HasColumn(&schema.SyncBlock{}, "has_frontend_move") {
		err = db.Migrator().AddColumn(&schema.SyncBlock{}, "has_frontend_move")
		if err != nil {
			return fmt.Errorf("failed to add has_frontend_move column to sync_blocks: %v", err)
		}

		// Set default value for newly added field
		err = db.Exec("UPDATE sync_blocks SET has_frontend_move = FALSE WHERE has_frontend_move IS NULL").Error
		if err != nil {
			return fmt.Errorf("failed to set default value for has_frontend_move: %v", err)
		}

		// Create index for new field
		err = db.Exec("CREATE INDEX idx_sync_blocks_has_frontend_move ON sync_blocks(has_frontend_move)").Error
		if err != nil {
			return fmt.Errorf("failed to create index on has_frontend_move: %v", err)
		}
	}

	// Check and add is_from_frontend field to sync_events table
	if !db.Migrator().HasColumn(&schema.SyncEvent{}, "is_from_frontend") {
		err = db.Migrator().AddColumn(&schema.SyncEvent{}, "is_from_frontend")
		if err != nil {
			return fmt.Errorf("failed to add is_from_frontend column to sync_events: %v", err)
		}

		// Set default value for newly added field
		err = db.Exec("UPDATE sync_events SET is_from_frontend = FALSE WHERE is_from_frontend IS NULL").Error
		if err != nil {
			return fmt.Errorf("failed to set default value for is_from_frontend: %v", err)
		}

		// Create index for new field
		err = db.Exec("CREATE INDEX idx_sync_events_is_from_frontend ON sync_events(is_from_frontend)").Error
		if err != nil {
			return fmt.Errorf("failed to create index on is_from_frontend: %v", err)
		}
	}

	return nil
}
