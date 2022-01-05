package service

import (
	"github.com/workfoxes/kayo/app/session"
	"github.com/workfoxes/kayo/internal/model"
)

type Order struct {
	Session *session.BaseSession
}

func (u *Order) GetOrders(page, limit int64) []*User {
	var users []*model.User
	u.Session.DB.Table("User").Find(&users).Limit(int(limit)).Offset(int(page * limit))
	return nil
}
