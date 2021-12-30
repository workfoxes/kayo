package model

import (
	"github.com/workfoxes/calypso/pkg/log"
	"gorm.io/gorm"
)

func AutoMigrateModel(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		&Group{},

		//&Strategy{},
		&ItemPointer{},
		&FilterCheck{},
		&IndicatorParams{},
	)
	if err != nil {
		log.S.Panic(err)
	}
}
