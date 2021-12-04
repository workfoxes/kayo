package model

import (
	"gorm.io/gorm"
)

type SystemModel struct {
	gorm.Model
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}
