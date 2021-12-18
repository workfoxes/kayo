package service

import (
	"github.com/workfoxes/kayo/app/session"
	"github.com/workfoxes/kayo/internal/model"
)

type User struct {
	Session *session.BaseSession
}

func (u *User) GetUsers(page, limit int64) []*model.User {
	var users []*model.User
	u.Session.DB.Find(&users).Limit(int(limit)).Offset(int(page * limit))
	return users
}

func (u *User) GetUser(userId uint) model.User {
	var user model.User
	u.Session.DB.Find(&user, userId)
	return user
}

func (u *User) UpdateUser(userId uint, user model.User) model.User {
	_user := u.GetUser(userId)
	u.Session.DB.Model(_user).Select("*").Updates(user)
	return _user
}

// CreateUser Create User Object in to the database based on the input object proved by the user
func (u *User) CreateUser(user model.User) model.User {
	u.Session.DB.Create(&user)
	return user
}

func (u *User) DeleteUser(userId uint) model.User {
	var user model.User
	u.Session.DB.Find(&user, userId)
	return user
}
