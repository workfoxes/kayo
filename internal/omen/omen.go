// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// omen is a controller which hold all the symbols of on going agents

package omen

import (
	"github.com/workfoxes/calypso/pkg/log"
	"github.com/workfoxes/kayo/internal/broker"
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator"
)

type Controller struct {
	Symbol     string
	Strategy   string
	IsLive     bool
	ItemChan   chan *common.Item
	Indicators []*indicator.TradeIndicator
	// Broker     *broker.StockBroker
}

var (
	ControllerMap map[string]*Controller
	BrokerMap     map[string]*broker.StockBroker
)

func GetBroker(_broker string) *broker.StockBroker {
	if BrokerMap != nil && BrokerMap[_broker] != nil {
		return BrokerMap[_broker]
	}
	brk := broker.NewBroker(_broker)
	BrokerMap[_broker] = brk
	return brk
}

// GetController will get Controller for symbol if exist else it will create new controller for the symbol
func GetController(symbol string) *Controller {
	if ControllerMap != nil && ControllerMap[symbol] != nil {
		return ControllerMap[symbol]
	}
	_controller := &Controller{
		Symbol:   symbol,
		ItemChan: make(chan *common.Item, 20),
	}
	ControllerMap[symbol] = _controller
	_broker := GetBroker(_controller.Symbol)
	(*_broker).Listen(_controller.Symbol, &_controller.ItemChan)
	go (*_controller).ProcessIndicator()
	return _controller
}

func (c *Controller) ProcessIndicator() {
	for cValue := range c.ItemChan {
		log.Info("ProcessIndicator", "symbol", c.Symbol, "value", cValue)
		for _, _indicator := range c.Indicators {
			go (*_indicator).Process(cValue)
		}
	}
}
