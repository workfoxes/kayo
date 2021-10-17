package broker

import (
	"github.com/workfoxes/kayo/internal/broker/binance"
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/broker/ib"
)

// NewBroker will create the broker instance based on the broker type that you are trying to work on
func NewBroker(brokerName string) *StockBroker {
	var broker StockBroker
	switch brokerName {
	case common.Binance:
		broker = new(binance.Binance)
	case common.InteractiveBroker:
		broker = new(ib.InteractiveBroker)
	}
	broker.Initialize()
	return &broker
}
