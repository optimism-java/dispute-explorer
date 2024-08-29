package v4

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/migration"
	"gorm.io/gorm"
)

var UpdateClaimDataPositionColumnTable = gormigrate.Migration{
	ID: "v4",
	Migrate: func(tx *gorm.DB) error {
		type GameClaimData struct {
			migration.Base
			Position string `json:"position" gorm:"type:varchar(128);notnull"`
		}
		return tx.AutoMigrate(&GameClaimData{})
	},
}
