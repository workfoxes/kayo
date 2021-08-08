package binance

const (
	ApiURL         = "https://api.binance.com"
	SpotApiURL     = "https://sapi.binance.com"
	CFuturesApiURL = "https://dapi.binance.com"
	UFuturesApiURL = "https://fapi.binance.com"

	StreamHostURL = "wss://stream.binance.com:9443"

	// Stream Endpoints

	RawStreamEndpoint      = "/ws"
	CombinedStreamEndpoint = "/stream?streams=%s"

	// Public endpoints

	ExchangeInfo      = "/api/v3/exchangeInfo"
	orderBookDepth    = "/api/v3/depth"
	recentTrades      = "/api/v3/trades"
	aggregatedTrades  = "/api/v3/aggTrades"
	candleStick       = "/api/v3/klines"
	averagePrice      = "/api/v3/avgPrice"
	priceChange       = "/api/v3/ticker/24hr"
	symbolPrice       = "/api/v3/ticker/price"
	bestPrice         = "/api/v3/ticker/bookTicker"
	userAccountStream = "/api/v3/userDataStream"
	perpExchangeInfo  = "/fapi/v1/exchangeInfo"
	historicalTrades  = "/api/v3/historicalTrades"

	// Authenticated endpoints
	newOrderTest      = "/api/v3/order/test"
	orderEndpoint     = "/api/v3/order"
	openOrders        = "/api/v3/openOrders"
	allOrders         = "/api/v3/allOrders"
	accountInfo       = "/api/v3/account"
	marginAccountInfo = "/sapi/v1/margin/account"
)
