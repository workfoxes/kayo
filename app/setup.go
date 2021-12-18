// this file will bootstrap the kayo tool with some basic configuration

package app

import (
	"github.com/workfoxes/kayo/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Bootstrap(DB *gorm.DB) {
	var s1 model.Strategy
	DB.Find(&s1).Preload(clause.Associations).Where("id = ?", 1).First(&s1)
	if s1.ID == 0 {
		s1 = model.Strategy{
			Name: "New Strategy",
			TakeProfit: model.ItemPointer{
				Type:  "Percentage",
				Value: 2,
			},
			StopLoss: model.ItemPointer{
				Type:  "Percentage",
				Value: -0.5,
			},
			BuyFilterCheck: []model.FilterCheck{
				{
					Indicator: "RSI",
					IndicatorParams: []model.IndicatorParams{
						{Key: "LookBackPeriod", Value: "14"},
						{Key: "AverageMethod", Value: "EMA"},
						{Key: "OverBuyLimit", Value: "80"},
						{Key: "OverSellLimit", Value: "30"},
						{Key: "CrossOver", Value: "true"},
					},
				},
			},
			SellFilterCheck: []model.FilterCheck{
				{
					Indicator: "RSI",
					IndicatorParams: []model.IndicatorParams{
						{Key: "LookBackPeriod", Value: "14"},
						{Key: "AverageMethod", Value: "EMA"},
						{Key: "OverBuyLimit", Value: "80"},
						{Key: "OverSellLimit", Value: "30"},
						{Key: "CrossOver", Value: "true"},
					},
				},
			},
		}
		DB.Create(&s1)
		DB.Commit()
	}
}
