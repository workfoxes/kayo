package model

import (
	"github.com/workfoxes/calypso/pkg/log"
	"gorm.io/gorm"
)

func AutoMigrateModel(db *gorm.DB) {
	err := db.AutoMigrate(
		&Order{},
	)
	if err != nil {
		log.Panic(err)
	}
}
