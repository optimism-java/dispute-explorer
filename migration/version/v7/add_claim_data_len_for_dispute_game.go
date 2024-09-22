package v7

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddClaimDataLenForDisputeGameTable = gormigrate.Migration{
	ID: "v7",
	Migrate: func(tx *gorm.DB) error {
		type DisputeGame struct {
			ClaimDataLen int64 `json:"claim_data_len" gorm:"type:bigint;notnull;index:dispute_claim_data_len_index;default:1"`
			GetLenStatus bool  `json:"get_len_status" gorm:"type:tinyint(1);notnull;default:0"`
		}
		return tx.Table("dispute_game").AutoMigrate(&DisputeGame{})
	},
	Rollback: func(db *gorm.DB) error {
		err := db.Migrator().DropColumn("dispute_game", "claim_data_len")
		if err != nil {
			return err
		}
		err = db.Migrator().DropColumn("dispute_game", "get_len_status")
		return err
	},
}
