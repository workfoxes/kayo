// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// omen is a controller which hold all the symbols of on going agents

package omen

import (
	"github.com/workfoxes/kayo/internal/broker"
)

type Controller struct {
	Symbol   string
	Strategy string
	IsLive   bool
	Broker   *broker.StockBroker
}

var ControllerMap map[string]*Controller

// GetController will get Controller for symbol if exist else it will create new controller for the symbol
func GetController(symbol string) *Controller {
	if ControllerMap != nil && ControllerMap[symbol] != nil {
		return ControllerMap[symbol]
	}
	return &Controller{
		Symbol: symbol,
	}
}
