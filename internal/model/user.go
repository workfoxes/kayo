package model

type UserType string

type User struct {
	Identity
	FirstName string   `gorm:"column:first_name;default:null" json:"FirstName"`
	LastName  string   `gorm:"column:last_name;default:null" json:"LastName"`
	Status    UserType `gorm:"column:status;default:Active" json:"Status"`
}
