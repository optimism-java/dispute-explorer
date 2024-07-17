package v2

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddCalculateLostForDisputeGameTable = gormigrate.Migration{
	ID: "v2",
	Migrate: func(tx *gorm.DB) error {
		type DisputeGame struct {
			CalculateLost bool `json:"calculate_lost" gorm:"type:tinyint(1);notnull;default:0"`
		}
		return tx.Table("dispute_game").AutoMigrate(&DisputeGame{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropColumn("dispute_game", "calculate_lost")
	},
}
