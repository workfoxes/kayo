package model

import (
	"gorm.io/gorm"
)

type SystemModel struct {
	gorm.Model
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}
