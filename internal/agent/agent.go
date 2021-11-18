// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// agent will controll all the agents in the system. It will be responsible for Holding the context for the order
package agent

type Agent struct {
	UserCtx string
}

func NewAgent() *Agent {
	return &Agent{
		UserCtx: "Vasanth",
	}
}
