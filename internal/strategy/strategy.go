package strategy

import (
	"github.com/workfoxes/kayo/internal/indicator"
	"github.com/workfoxes/kayo/internal/model"
)

type Strategy struct {
	model.SystemModel
	Name            string              `gorm:"column:Name;" json:"Name"`
	Description     string              `gorm:"column:Description;default:null" json:"Description"`
	TakeProfit      model.ItemPointer   `json:"TakeProfit,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	StopLoss        model.ItemPointer   `json:"StopLoss,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	BuyFilterCheck  []model.FilterCheck `json:"BuyFilterCheck,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	SellFilterCheck []model.FilterCheck `json:"SellFilterCheck,omitempty" gorm:"foreignKey:StrategyID;references:id"`

	BuyFilters  []*indicator.TradeIndicator
	SellFilters []*indicator.TradeIndicator
}

//type ItemPointer struct {
//	model.SystemModel
//	Type       string  `json:"Type"`
//	Value      float64 `json:"Value"`
//	StrategyID int     `json:"StrategyID"`
//}
//type FilterCheck struct {
//	model.SystemModel
//	Indicator       string            `json:"Indicator"`
//	IndicatorParams []IndicatorParams `json:"IndicatorParams,omitempty" gorm:"foreignKey:CheckerID;references:id"`
//	StrategyID      int               `gorm:"column:StrategyID;" json:"StrategyID"`
//}
//type IndicatorParams struct {
//	model.SystemModel
//	Key       string `gorm:"column:Key;" json:"Key"`
//	Value     string `gorm:"column:Value;" json:"Value"`
//	CheckerID int    `json:"CheckerID"`
//}

// ParseStrategy will parse the Strategy, so it will be ready to process
func ParseStrategy(strategy *Strategy) {
	for _, buyFilter := range strategy.BuyFilterCheck {
		_indicator := indicator.NewIndicator(buyFilter.Indicator)
		strategy.BuyFilters = append(strategy.BuyFilters, _indicator)
		(*_indicator).Initialize()
	}
	for _, sellFilter := range strategy.SellFilterCheck {
		_indicator := indicator.NewIndicator(sellFilter.Indicator)
		strategy.SellFilters = append(strategy.SellFilters, _indicator)
	}
}

func (s *Strategy) GetIndicators() []*indicator.TradeIndicator {
	var _i []*indicator.TradeIndicator
	_i = append(_i, s.BuyFilters...)
	_i = append(_i, s.SellFilters...)
	return _i
}
