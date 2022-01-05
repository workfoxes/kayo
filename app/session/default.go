package session

import "gorm.io/gorm"

type BaseSession struct {
	DB *gorm.DB
	TX *gorm.DB
}

func find(id string) *BaseSession {
	return nil
}
