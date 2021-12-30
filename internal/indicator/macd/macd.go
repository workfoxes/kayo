package macd

import (
	"github.com/workfoxes/calypso/pkg/dsa"
	"github.com/workfoxes/calypso/pkg/log"
	commonB "github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator/common"
	"github.com/workfoxes/kayo/internal/model"
	"strconv"
)

type MACD struct {
	common.BaseIndicator
	LookBackPeriod      int
	AverageMethod       string
	BuyLimit, SellLimit float64
	CrossOver           bool
	RecentClose         *dsa.CList
	RSI                 *dsa.CList
	IsComputable        bool
}

func (macd *MACD) Process(item *commonB.Item) bool {
	macd.RecentClose.Enqueue(item.ClosePrice)
	if macd.IsComputable {
		_ = macd.RecentClose.Q[(len(macd.RecentClose.Q) - macd.LookBackPeriod):]
	} else if len(macd.RecentClose.Q) > macd.LookBackPeriod {
		macd.IsComputable = true
	}
	return false
}

func (macd *MACD) SetFilterCheck(check model.FilterCheck) error {
	macd.Name = check.Indicator
	var err error
	for _, param := range check.IndicatorParams {
		switch param.Key {
		case "LookBackPeriod":
			i, err := strconv.Atoi(param.Value)
			if err != nil {
				return err
			}
			macd.LookBackPeriod = i
		case "AverageMethod":
			macd.AverageMethod = param.Value
		case "OverBuyLimit":
			macd.BuyLimit, err = strconv.ParseFloat(param.Value, 64)
			if err != nil {
				return err
			}
		case "OverSellLimit":
			macd.SellLimit, err = strconv.ParseFloat(param.Value, 64)
			if err != nil {
				return err
			}
		default:
			log.S.Info("Unknown parameter to load into rsi ", param)
		}
	}
	macd.RecentClose = dsa.NewCList(macd.LookBackPeriod * 2)
	macd.RSI = dsa.NewCList(macd.LookBackPeriod * 2)
	return nil
}
