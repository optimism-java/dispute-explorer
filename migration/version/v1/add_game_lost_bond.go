package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/optimism-java/dispute-explorer/migration"
	"gorm.io/gorm"
)

var AddGameLostBondTable = gormigrate.Migration{
	ID: "v1",
	Migrate: func(tx *gorm.DB) error {
		type GameLostBond struct {
			migration.Base
			GameContract string `json:"game_contract" gorm:"type:varchar(42);notnull"`
			Address      string `json:"address" gorm:"type:varchar(64);notnull"`
			Bond         string `json:"bond" gorm:"type:bigint;notnull"`
		}
		return tx.AutoMigrate(&GameLostBond{})
	},
	Rollback: func(db *gorm.DB) error {
		return db.Migrator().DropTable("game_lost_bond")
	},
}
