package main

import (
	"flag"
	"strings"
	"time"

	"github.com/workfoxes/kayo/pkg/runner"
)

var (
	_broker    = flag.String("broker", "Binance", "Broker : that will be used for trading")
	_symbol    = flag.String("symbol", "BTCUSDT", "Symbol : which will tracked and traded by bot")
	_isLive    = flag.Bool("live", false, "live : Is trading on the live platform")
	_strategy  = flag.String("strategy", "default", "strategy : will use to stick with the specific strategy")
	_indicator = flag.String("indicator", "RSI|MACD", "indicator : will use to stick with the specific indicator")
)

func main() {
	_symbols := strings.SplitAfter(*_symbol, ",")
	for _, value := range _symbols {
		_runner := runner.Runner{BrokerName: *_broker, Symbol: value,
			IsLive:    *_isLive,
			Strategy:  *_strategy,
			Indicator: *_indicator}
		_runner.Initialize()
	}
	for {
		time.Sleep(time.Second * 1000000)
	}
}
