// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package agent will control all the agents in the system. It will be responsible for Holding the context for the order
package agent

import (
	"github.com/workfoxes/calypso/pkg/log"
	"github.com/workfoxes/kayo/internal/model"
	"github.com/workfoxes/kayo/internal/strategy"
)

type Agent struct {
	Symbol   string
	UserCtx  *model.User
	Strategy strategy.Strategy
}

func (a *Agent) Run() {
	log.Info("Agent is running")
}

func (a *Agent) Stop() {
	log.Info("Stopping Agent")
}
