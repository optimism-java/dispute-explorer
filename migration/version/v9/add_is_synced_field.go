package v9

import (
	"fmt"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"gorm.io/gorm"
)

var AddIsSyncedFieldTable = gormigrate.Migration{
	ID:      "20240829_add_is_synced_field",
	Migrate: AddIsSyncedToFrontendMoveTransactions,
}

func AddIsSyncedToFrontendMoveTransactions(db *gorm.DB) error {
	// Check and add is_synced field to frontend_move_transactions table
	if !db.Migrator().HasColumn(&schema.FrontendMoveTransaction{}, "is_synced") {
		err := db.Migrator().AddColumn(&schema.FrontendMoveTransaction{}, "is_synced")
		if err != nil {
			return fmt.Errorf("failed to add is_synced column to frontend_move_transactions: %v", err)
		}

		// Set default value for newly added field
		err = db.Model(&schema.FrontendMoveTransaction{}).Update("is_synced", false).Error
		if err != nil {
			return fmt.Errorf("failed to set default value for is_synced: %v", err)
		}
	}

	return nil
}
