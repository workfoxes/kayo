package indicator

import (
	"github.com/workfoxes/kayo/internal/indicator/adl"
	"github.com/workfoxes/kayo/internal/indicator/adx"
	"github.com/workfoxes/kayo/internal/indicator/macd"
	"github.com/workfoxes/kayo/internal/indicator/obv"
	"github.com/workfoxes/kayo/internal/indicator/rsi"
	"github.com/workfoxes/kayo/internal/indicator/so"
)

// SupportedTradingIndicators  - list all supporting trading indicator

const (
	RSI                          = "RSI"
	MACD                         = "MACD"
	StochasticOscillator         = "StochasticOscillator"
	AverageDirectionalIndex      = "AverageDirectionalIndex"
	OnBalanceVolume              = "OnBalanceVolume"
	AccumulationDistributionLine = "AccumulationDistributionLine"
)

type BaseIndicator struct {
}

// NewIndicator : will create new indicator based on the item pattern
func NewIndicator(name string) TradeIndicator {
	var indicator TradeIndicator
	switch name {
	case RSI:
		indicator = new(rsi.Indicator)
	case MACD:
		indicator = new(macd.Indicator)
	case StochasticOscillator:
		indicator = new(so.Indicator)
	case AverageDirectionalIndex:
		indicator = new(adx.Indicator)
	case OnBalanceVolume:
		indicator = new(obv.Indicator)
	case AccumulationDistributionLine:
		indicator = new(adl.Indicator)
	}
	return indicator
}
