package model

import "github.com/workfoxes/kayo/internal/model"

type Order struct {
	model.SystemModel
	ID     int    `gorm:"column:ID;primary_key:auto_increment" json:"ID"`
	Symbol string `gorm:"column:Name;default:null" json:"Symbol"`
	Email  string `gorm:"size:255;column:email;index:idx_email,unique"`
}
