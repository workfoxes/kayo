package model

type Identity struct {
	SystemModel
	Name  string `gorm:"column:name;default:null" json:"Name"`
	Email string `gorm:"size:255;column:email;index:idx_email,unique"`
}
