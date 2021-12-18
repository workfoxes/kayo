package rsi

import (
	commonB "github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator/common"
)

// RSI : Relative Strength Index (RSI).
// Compares the magnitude of recent gains and losses over a specified time
// period to measure speed and change of price movements of a security. It is
// primarily used to attempt to identify overbought or oversold conditions in
// the trading of an asset.
// See [Relative Strength Index (RSI)](https://www.investopedia.com/terms/r/rsi.asp).
type RSI struct {
	common.BaseIndicator
	LookBackPeriod int
	AverageMethod  string
}

type Params struct {
	LookBackPeriod int
	AverageMethod  string
}

func (i *RSI) Process(item *commonB.Item) {
	i.Name = "Name"
	i.ShortName = "ShortName"
}
