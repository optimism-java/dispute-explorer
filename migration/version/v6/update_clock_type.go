package v6

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var UpdateClaimDataClockColumnTable = gormigrate.Migration{
	ID: "v6",
	Migrate: func(tx *gorm.DB) error {
		type GameClaimData struct {
			Clock string `json:"clock" gorm:"type:varchar(128);notnull"`
		}
		return tx.AutoMigrate(&GameClaimData{})
	},
}
