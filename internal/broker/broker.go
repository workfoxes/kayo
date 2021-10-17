package broker

import (
	"github.com/workfoxes/kayo/internal/broker/binance"
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/broker/ib"
	"github.com/workfoxes/kayo/internal/errors"
)

// NewBroker will create the broker instance based on the broker type that you are trying to work on
// this function follows the factory pattern to get the Broker object based on the args
// 	params : brokerName will string value that will be name of the broker will be used for trading
func NewBroker(brokerName string) *StockBroker {
	var broker StockBroker
	switch brokerName {
	case common.Binance:
		broker = new(binance.Binance)
	case common.InteractiveBroker:
		broker = new(ib.InteractiveBroker)
	default:
		panic(errors.IndicatorNotFound)
	}
	broker.Initialize()
	return &broker
}
