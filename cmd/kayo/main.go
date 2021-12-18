package main

import (
	"flag"
	"github.com/workfoxes/calypso/pkg/log"
	"time"
)

var (
	_broker    = flag.String("broker", "Binance", "Broker : that will be used for trading")
	_symbol    = flag.String("symbol", "BTCUSDT", "Symbol : which will tracked and traded by bot")
	_isLive    = flag.Bool("live", false, "live : Is trading on the live platform")
	_strategy  = flag.String("strategy", "default", "strategy : will use to stick with the specific strategy")
	_indicator = flag.String("indicator", "RSI|MACD", "indicator : will use to stick with the specific indicator")
)

func main() {
	//go func() {
	//	app.StartKayoServer()
	//}()
	log.Info("Starting the Bot")
	//_symbols := strings.SplitAfter(*_symbol, ",")
	//for _, value := range _symbols {
	//	go func(value string) {
	//		_agent := &agent.Agent{
	//			Symbol:   value,
	//			UserCtx:  &model.User{},
	//			Strategy: s1,
	//		}
	//		_agent.Run()
	//	}(value)
	//}
	for {
		time.Sleep(time.Second * 1000000)
	}
}
