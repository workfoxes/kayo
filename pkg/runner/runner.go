package runner

import (
	"github.com/workfoxes/kayo/internal/broker"
	"github.com/workfoxes/kayo/internal/broker/common"
)

// Runner : Holds the current execution state symbol and strategy
type Runner struct {
	Symbol, Strategy, BrokerName string
	IsLive                       bool
	Broker                       *broker.StockBroker
}

func (r *Runner) Initialize() {
	listener := make(chan *common.Item, 20)
	r.Broker = broker.NewBroker(r.BrokerName)
	(*r.Broker).Listen(r.Symbol, listener)
}
