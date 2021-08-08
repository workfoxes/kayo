package broker

// StockBroker enforces standard functions for all exchanges supported in kayo
type StockBroker interface {
	Initialize()
	RegisterWebsocketClient()
}
