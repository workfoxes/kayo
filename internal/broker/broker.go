package broker

import (
	"github.com/workfoxes/kayo/internal/broker/binance"
	"github.com/workfoxes/kayo/internal/broker/default"
	"github.com/workfoxes/kayo/internal/broker/ib"
)

// NewBroker will create the broker instance based on the broker type that you are trying to work on
func NewBroker(brokerName string) StockBroker {
	var broker StockBroker
	switch brokerName {
	case _default.Binance:
		broker = new(binance.Binance)
	case _default.InteractiveBroker:
		broker = new(ib.InteractiveBroker)
	}
	broker.Initialize()
	return broker
}
