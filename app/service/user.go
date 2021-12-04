package service

import (
	"github.com/workfoxes/kayo/app/session"
	"github.com/workfoxes/kayo/internal/model"
)

type User struct {
	Session *session.BaseSession
}

func (u *User) GetUsers(page, limit int64) []*User {
	var users []*model.User
	u.Session.DB.Table("User").Find(&users).Limit(int(limit)).Offset(int(page * limit))
	return nil
}
