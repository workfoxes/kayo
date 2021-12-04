package runner

import (
	"github.com/workfoxes/calypso/pkg/log"
	"strings"

	"github.com/workfoxes/kayo/internal/broker"
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator"
)

// Runner : Holds the current execution state symbol and strategy
type Runner struct {
	Symbol, Strategy, BrokerName, Indicator string
	IsLive                                  bool
	Broker                                  *broker.StockBroker
	Indicators                              []*indicator.TradeIndicator
}

func (r *Runner) Initialize() {
	listener := make(chan *common.Item, 20)
	r.Broker = broker.NewBroker(r.BrokerName)
	(*r.Broker).Listen(r.Symbol, &listener)

	var _indicators []string
	_indicators = append(_indicators, strings.Split(r.Indicator, "|")...)
	for _, _indicator := range _indicators {
		r.Indicators = append(r.Indicators, indicator.NewIndicator(_indicator))
	}
	r.checkLine(&listener)
}

func (r *Runner) checkLine(itemChan *chan *common.Item) {
	for cValue := range *itemChan {
		log.Debug(cValue)
	}
}
