package ib

import (
	"github.com/workfoxes/kayo/internal/broker/common"
)

type InteractiveBroker struct {
	common.BaseBroker
}

func (b *InteractiveBroker) Initialize() {
	b.Name = common.InteractiveBroker
	b.IsWebSocketSupported = true
}
