package ib

import (
	"github.com/workfoxes/kayo/internal/broker/default"
)

type InteractiveBroker struct {
	_default.BaseBroker
}

func (b *InteractiveBroker) Initialize() {
	b.Name = _default.InteractiveBroker
	b.IsWebSocketSupported = true
}

func (b *InteractiveBroker) RegisterWebsocketClient() {
}
