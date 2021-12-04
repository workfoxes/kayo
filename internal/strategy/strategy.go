package strategy

import (
	"github.com/workfoxes/kayo/internal/indicator"
)

type Strategy struct {
	ID              string        `json:"_id"`
	Strategy        string        `json:"Strategy"`
	TakeProfit      ItemPointer   `json:"TakeProfit"`
	StopLoss        ItemPointer   `json:"StopLoss"`
	BuyFilterCheck  []FilterCheck `json:"BuyFilterCheck"`
	SellFilterCheck []FilterCheck `json:"SellFilterCheck"`

	BuyFilters  []*indicator.TradeIndicator `json:"BuyFilters"`
	SellFilters []*indicator.TradeIndicator `json:"SellFilters"`
}
type ItemPointer struct {
	Type  string  `json:"Type"`
	Value float64 `json:"Value"`
}
type IndicatorParams struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
type FilterCheck struct {
	Indicator       string            `json:"Indicator"`
	IndicatorParams []IndicatorParams `json:"IndicatorParams"`
}

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
