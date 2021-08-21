package runner

import (
	"github.com/workfoxes/kayo/internal/broker"
	_default "github.com/workfoxes/kayo/internal/broker/default"
)

// Runner : Holds the current execution state symbol and strategy
type Runner struct {
	Symbol     string
	Strategy   string
	BrokerName string
	IsLive     bool
	Broker     broker.StockBroker
}

func (r *Runner) Initialize() {
	listener := make(chan *_default.Item, 2)
	r.Broker = broker.NewBroker(r.BrokerName)
	r.Broker.Listen(listener)
}
