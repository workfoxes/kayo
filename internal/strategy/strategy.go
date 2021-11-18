package strategy

import (
	"fmt"

	"github.com/workfoxes/kayo/internal/indicator"
)

type Strategy struct {
	ID              string        `json:"_id"`
	Strategy        string        `json:"Strategy"`
	LimitOnClose    ItemPointer   `json:"LimitOnClose"`
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

func ParseStragety(strategy *Strategy) {
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

func Strategy001() {
	s1 := &Strategy{
		ID:       "001",
		Strategy: "001",
		LimitOnClose: ItemPointer{
			Type:  "Percentage",
			Value: 2,
		},
		StopLoss: ItemPointer{
			Type:  "Percentage",
			Value: -0.5,
		},
		BuyFilterCheck: []FilterCheck{
			{
				Indicator: "RSI",
				IndicatorParams: []IndicatorParams{
					{Key: "LookBackPeriod", Value: "14"},
					{Key: "AverageMethod", Value: "EMA"},
					{Key: "OverBuyLimit", Value: "80"},
					{Key: "OverSellLimit", Value: "30"},
					{Key: "CrossOver", Value: "true"},
				},
			},
		},
		SellFilterCheck: []FilterCheck{
			{
				Indicator: "RSI",
				IndicatorParams: []IndicatorParams{
					{Key: "LookBackPeriod", Value: "14"},
					{Key: "AverageMethod", Value: "EMA"},
					{Key: "OverBuyLimit", Value: "80"},
					{Key: "OverSellLimit", Value: "30"},
					{Key: "CrossOver", Value: "true"},
				},
			},
		},
	}
	fmt.Println(s1)

}
