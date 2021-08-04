package main

import (
	"flag"
)

var symbols []string

var (
	service = flag.String("service", "Binance", "Service that will be used for trading")
	symbol = flag.String("symbol", "BTCUSDT", "Symbol which will tracked and traded by bot")
	api_key = flag.String("key", "", "the path to watch")
	isLive = flag.Bool("live", false, "live : Is trading on the live platform")
)

// main : Main will start the cron application for Kayo
func main() {
	flag.Parse()

}