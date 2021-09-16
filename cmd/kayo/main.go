package main

import (
	"flag"
	"github.com/workfoxes/kayo/pkg/runner"
	"strings"
	"time"
)

var (
	_broker   = flag.String("broker", "Binance", "Broker : that will be used for trading")
	_symbol   = flag.String("symbol", "BTCUSDT", "Symbol : which will tracked and traded by bot")
	_isLive   = flag.Bool("live", false, "live : Is trading on the live platform")
	_strategy = flag.String("strategy", "default", "strategy : will use to stick with the specific strategy")
)

func main() {
	_symbols := strings.SplitAfter(*_symbol, ",")
	for _, value := range _symbols {
		_runner := runner.Runner{BrokerName: *_broker, Symbol: value,
			IsLive:   *_isLive,
			Strategy: *_strategy}
		_runner.Initialize()
	}
	for {
		time.Sleep(time.Second * 1000000)
	}
}
