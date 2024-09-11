package v5

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddOnChainStatusForDisputeGameTable = gormigrate.Migration{
	ID: "v5",
	Migrate: func(tx *gorm.DB) error {
		type DisputeGame struct {
			OnChainStatus string `json:"on_chain_status" gorm:"type:varchar(32);notnull;index:dispute_on_chain_status_index;default:valid"`
		}
		type GameClaimData struct {
			OnChainStatus string `json:"on_chain_status" gorm:"type:varchar(32);notnull;index:claim_on_chain_status_index;default:valid"`
		}
		err := tx.Table("dispute_game").AutoMigrate(&DisputeGame{})
		if err != nil {
			return err
		}
		return tx.AutoMigrate(&GameClaimData{})
	},
	Rollback: func(db *gorm.DB) error {
		err := db.Migrator().DropColumn("dispute_game", "on_chain_status")
		if err != nil {
			return err
		}
		err = db.Migrator().DropColumn("game_claim_data", "on_chain_status")
		return err
	},
}
