package main

import (
	"flag"
	"github.com/workfoxes/kayo/internal/broker"
	"strings"
)

type KayoWorker struct {
	Name       string
	BrokerName string
	Symbol     []string
	IsLive     bool
	Strategy   string
	Broker     broker.StockBroker
}

var (
	_broker   = flag.String("broker", "Binance", "Broker : that will be used for trading")
	_symbol   = flag.String("symbol", "BTCUSDT", "Symbol : which will tracked and traded by bot")
	_isLive   = flag.Bool("live", false, "live : Is trading on the live platform")
	_strategy = flag.String("strategy", "default", "strategy : will use to stick with the specific strategy")
)

func main() {
	_symbols := strings.SplitAfter(*_symbol, ",")
	kayo := &KayoWorker{
		Name:       "Kayo",
		BrokerName: *_broker,
		Symbol:     _symbols,
		IsLive:     *_isLive,
		Strategy:   *_strategy,
	}
	kayo.Broker = *broker.NewBroker(kayo.BrokerName)
}
