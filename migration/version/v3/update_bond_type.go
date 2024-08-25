package v3

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/migration"
	"gorm.io/gorm"
)

var UpdateLostBondAndClaimDataTable = gormigrate.Migration{
	ID: "v3",
	Migrate: func(tx *gorm.DB) error {
		type GameLostBond struct {
			migration.Base
			Bond string `json:"bond" gorm:"type:varchar(128);notnull"`
		}
		type GameClaimData struct {
			migration.Base
			Bond string `json:"bond" gorm:"type:varchar(128);notnull"`
		}
		return tx.AutoMigrate(&GameLostBond{}, &GameClaimData{})
	},
	Rollback: func(db *gorm.DB) error {
		return fmt.Errorf("update column type error")
	},
}
