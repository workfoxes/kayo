package main

import (
	"flag"
	"strings"
	"time"

	"github.com/workfoxes/kayo/app"

	"github.com/workfoxes/kayo/internal/agent"
	"github.com/workfoxes/kayo/internal/model"
	"github.com/workfoxes/kayo/internal/strategy"
)

var (
	_broker    = flag.String("broker", "Binance", "Broker : that will be used for trading")
	_symbol    = flag.String("symbol", "BTCUSDT", "Symbol : which will tracked and traded by bot")
	_isLive    = flag.Bool("live", false, "live : Is trading on the live platform")
	_strategy  = flag.String("strategy", "default", "strategy : will use to stick with the specific strategy")
	_indicator = flag.String("indicator", "RSI|MACD", "indicator : will use to stick with the specific indicator")
)

func main() {
	go func() {
		app.StartKayoServer()
	}()
	s1 := strategy.Strategy{
		ID:       "001",
		Strategy: "001",
		TakeProfit: strategy.ItemPointer{
			Type:  "Percentage",
			Value: 2,
		},
		StopLoss: strategy.ItemPointer{
			Type:  "Percentage",
			Value: -0.5,
		},
		BuyFilterCheck: []strategy.FilterCheck{
			{
				Indicator: "RSI",
				IndicatorParams: []strategy.IndicatorParams{
					{Key: "LookBackPeriod", Value: "14"},
					{Key: "AverageMethod", Value: "EMA"},
					{Key: "OverBuyLimit", Value: "80"},
					{Key: "OverSellLimit", Value: "30"},
					{Key: "CrossOver", Value: "true"},
				},
			},
		},
		SellFilterCheck: []strategy.FilterCheck{
			{
				Indicator: "RSI",
				IndicatorParams: []strategy.IndicatorParams{
					{Key: "LookBackPeriod", Value: "14"},
					{Key: "AverageMethod", Value: "EMA"},
					{Key: "OverBuyLimit", Value: "80"},
					{Key: "OverSellLimit", Value: "30"},
					{Key: "CrossOver", Value: "true"},
				},
			},
		},
	}
	_symbols := strings.SplitAfter(*_symbol, ",")
	for _, value := range _symbols {
		go func(value string) {
			_agent := &agent.Agent{
				Symbol:   value,
				UserCtx:  &model.User{},
				Strategy: s1,
			}
			_agent.Run()
		}(value)
	}
	for {
		time.Sleep(time.Second * 1000000)
	}
}
