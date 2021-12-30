package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/log"
	"github.com/workfoxes/kayo/internal/utils"
	"github.com/workfoxes/tripwire/binance"
	"time"
)

func main() {
	c := binance.NewWebClient(false)
	_ch, _ := c.KlineWebsocket("bnbbtc", "1m")
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	option := influxdb2.DefaultOptions().SetLogLevel(log.DebugLevel)
	option.WriteOptions().SetBatchSize(10)
	client := influxdb2.NewClientWithOptions("http://localhost:8086",
		"u6k1U8KNjZ-su0yl9hUVAjP3GMtdTRpeITnTAsJ_RXtjC19iYKr5Ldyw5gc2hIdJGtFNzfq3P4D0rXUrDZtI2Q==", option)
	// always close client at the end
	defer client.Close()
	// get non-blocking write client
	writeAPI := client.WriteAPI("kayo", "kayo")

	defer writeAPI.Flush()

	go func() {
		select {
		default:
			for cValue := range _ch {
				c.Info(cValue)
				p := influxdb2.NewPointWithMeasurement(cValue.Event.Symbol).
					AddTag("symbol", cValue.Event.Symbol).
					AddField("high", utils.ParseFloat(cValue.Kline.High)).
					AddField("low", utils.ParseFloat(cValue.Kline.Low)).
					AddField("open", utils.ParseFloat(cValue.Kline.Open)).
					AddField("close", utils.ParseFloat(cValue.Kline.Close)).
					AddField("volume", utils.ParseFloat(cValue.Kline.Volume)).
					AddField("low", cValue.Kline.NumberOfTrades).
					SetTime(time.Now())

				//
				//p := influxdb2.NewPointWithMeasurement("stat").
				//	AddTag("unit", "temperature").
				//	AddField("avg", 23.2).
				//	AddField("max", 45).
				//	SetTime(time.Now())

				// write point asynchronously
				writeAPI.WritePoint(p)
				c.Info(p)
				if cValue.Final {
					c.Info("Final")
				}
			}
		}

	}()
	for {
		time.Sleep(time.Second * 1000)
	}
}
