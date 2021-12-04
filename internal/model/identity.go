package model

type Identity struct {
	SystemModel
	ID    int    `gorm:"column:ID;primary_key:auto_increment" json:"ID"`
	Name  string `gorm:"column:Name;default:null" json:"Name"`
	Email string `gorm:"size:255;column:email;index:idx_email,unique"`
}
