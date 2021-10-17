package broker

import "github.com/workfoxes/kayo/internal/broker/common"

// StockBroker enforces standard functions for all exchanges supported in kayo
type StockBroker interface {
	// Initialize : will initialize the broker config and setup
	Initialize()
	RegisterWebsocketClient(url string)
	// Listen : will listen the market data changes for the selected symbol
	Listen(symbol string, itemChan *chan *common.Item)
}
