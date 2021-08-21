package broker

import _default "github.com/workfoxes/kayo/internal/broker/default"

// StockBroker enforces standard functions for all exchanges supported in kayo
type StockBroker interface {
	Initialize()
	RegisterWebsocketClient(url string)
	Listen(itemChan chan *_default.Item)
}
