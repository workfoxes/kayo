package rsi

import (
	"github.com/workfoxes/calypso/pkg/dsa"
	"github.com/workfoxes/calypso/pkg/log"
	commonB "github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator/common"
	"github.com/workfoxes/kayo/internal/model"
	"strconv"
)

// RSI : Relative Strength Index (RSI).
// Compares the magnitude of recent gains and losses over a specified time
// period to measure speed and change of price movements of a security. It is
// primarily used to attempt to identify overbought or oversold conditions in
// the trading of an asset.
// See [Relative Strength Index (RSI)](https://www.investopedia.com/terms/r/rsi.asp).
type RSI struct {
	common.BaseIndicator
	LookBackPeriod      int
	AverageMethod       string
	BuyLimit, SellLimit float64
	CrossOver           bool
	RecentClose         *dsa.CList
	RSI                 *dsa.CList
	IsComputable        bool
}

func (rsi *RSI) Process(item *commonB.Item) bool {
	rsi.RecentClose.Enqueue(item.ClosePrice)
	if rsi.IsComputable {
		_ = rsi.RecentClose.Q[(len(rsi.RecentClose.Q) - rsi.LookBackPeriod):]
		//ta.Rsi(rsiFloat, rsi.LookBackPeriod)
	} else if len(rsi.RecentClose.Q) > rsi.LookBackPeriod {
		rsi.IsComputable = true
	}
	item.IndicatorStatus[rsi.ID()] = false
	return false
}

func (rsi *RSI) SetFilterCheck(check model.FilterCheck) error {
	rsi.Name = check.Indicator
	var err error
	for _, param := range check.IndicatorParams {
		switch param.Key {
		case "LookBackPeriod":
			i, err := strconv.Atoi(param.Value)
			if err != nil {
				return err
			}
			rsi.LookBackPeriod = i
		case "AverageMethod":
			rsi.AverageMethod = param.Value
		case "OverBuyLimit":
			rsi.BuyLimit, err = strconv.ParseFloat(param.Value, 64)
			if err != nil {
				return err
			}
		case "OverSellLimit":
			rsi.SellLimit, err = strconv.ParseFloat(param.Value, 64)
			if err != nil {
				return err
			}
		case "CrossOver":
			rsi.CrossOver = param.Value == "true"
		default:
			log.S.Info("Unknown parameter to load into rsi  - ", param.Key, param)
		}
	}
	rsi.RecentClose = dsa.NewCList(rsi.LookBackPeriod * 2)
	rsi.RSI = dsa.NewCList(rsi.LookBackPeriod * 2)
	return nil
}
