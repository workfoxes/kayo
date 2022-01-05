package model

type Group struct {
	Identity
	Description string `gorm:"column:description;default:null" json:"Description"`
}
