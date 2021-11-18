package indicator

import (
	"github.com/workfoxes/kayo/internal/errors"
	"github.com/workfoxes/kayo/internal/indicator/adl"
	"github.com/workfoxes/kayo/internal/indicator/adx"
	"github.com/workfoxes/kayo/internal/indicator/common"
	"github.com/workfoxes/kayo/internal/indicator/macd"
	"github.com/workfoxes/kayo/internal/indicator/obv"
	"github.com/workfoxes/kayo/internal/indicator/rsi"
	"github.com/workfoxes/kayo/internal/indicator/so"
)

// NewIndicator : will create new indicator based on the item pattern
func NewIndicator(name string) *TradeIndicator {
	var indicator TradeIndicator
	switch name {
	case common.RSI:
		indicator = new(rsi.RSI)
	case common.MACD:
		indicator = new(macd.Indicator)
	case common.StochasticOscillator:
		indicator = new(so.Indicator)
	case common.AverageDirectionalIndex:
		indicator = new(adx.Indicator)
	case common.OnBalanceVolume:
		indicator = new(obv.Indicator)
	case common.AccumulationDistributionLine:
		indicator = new(adl.Indicator)
	default:
		panic(errors.IndicatorNotFound)
	}
	return &indicator
}
