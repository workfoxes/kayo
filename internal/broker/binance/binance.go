package binance

import (
	"github.com/workfoxes/kayo/internal/broker/default"
)

type Binance struct {
	_default.BaseBroker
}

func (b *Binance) Initialize() {
	b.Name = _default.Binance
	b.IsWebSocketSupported = true
}

func (b *Binance) RegisterWebsocketClient() {

}
