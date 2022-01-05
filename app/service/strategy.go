package service

import (
	"github.com/workfoxes/kayo/app/session"
	"github.com/workfoxes/kayo/internal/model"
	"gorm.io/gorm/clause"
)

type Strategy struct {
	Session *session.BaseSession
}

func (u *Strategy) GetStrategies(page, limit int64) []*model.Strategy {
	var strategy []*model.Strategy
	u.Session.DB.Find(&strategy).Limit(int(limit)).Offset(int(page * limit))
	return strategy
}

func (u *Strategy) GetStrategy(id uint) model.Strategy {
	var strategy *model.Strategy
	u.Session.DB.First(&strategy, id).Preload(clause.Associations)
	return *strategy
}

func (u *Strategy) CreateStrategy(page, limit int64) []*model.Strategy {
	var strategy []*model.Strategy
	u.Session.DB.Find(&strategy).Limit(int(limit)).Offset(int(page * limit))
	return strategy
}

func (u *Strategy) UpdateStrategy(page, limit int64) []*model.Strategy {
	var strategy []*model.Strategy
	u.Session.DB.Find(&strategy).Limit(int(limit)).Offset(int(page * limit))
	return strategy
}
