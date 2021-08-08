package main

import (
	"flag"
	"github.com/go-redis/redis/v8"
	"strings"
)

type KayoWorker struct {
	Name               string
	Broker             string
	Symbol             []string
	IsLive             bool
	DefaultChannelName string
	DefaultChannel     *redis.PubSub
	Client             *redis.Client
}

var (
	broker = flag.String("broker", "Binance", "Broker that will be used for trading")
	symbol = flag.String("symbol", "BTCUSDT", "Symbol which will tracked and traded by bot")
	isLive = flag.Bool("live", false, "live : Is trading on the live platform")
)

func main() {
	_symbols := strings.SplitAfter(*symbol, ",")
	_ = &KayoWorker{
		Name:   "Kayo",
		Broker: *broker,
		Symbol: _symbols,
		IsLive: *isLive,
	}
}
