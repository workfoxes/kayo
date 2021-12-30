package main

import (
	"flag"
	"fmt"
	"github.com/workfoxes/calypso/pkg/log"
	"github.com/workfoxes/kayo/internal/broker"
	"github.com/workfoxes/kayo/pkg/constant"
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
	BootKayo()
}

// BootKayo : will bootstrap the bot check for all the running workers and start them
func BootKayo() {
	fmt.Printf(constant.ASCIILogo)
	log.S.Info("Starting the Kayo Bot")
}
