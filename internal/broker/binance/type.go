package binance

type WebSocketRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

type WebSocketResponse struct {
	EventType string `json:"e"`
	EventName int64  `json:"E"`
	Symbol    string `json:"s"`
	KPI       Kpi    `json:"k"`
}

type Kpi struct {
	KStartTime       int64  `json:"t"`
	KSEndTime        int64  `json:"T"`
	Symbol           string `json:"s"`
	Interval         string `json:"i"`
	FirstTradeID     int    `json:"f"`
	LastTradeID      int    `json:"L"`
	OpenPrice        string `json:"o"`
	ClosePrice       string `json:"c"`
	HighestPrice     string `json:"h"`
	LowestPrice      string `json:"l"`
	BassAssetVolume  string `json:"v"`
	NoOfTrade        int    `json:"n"`
	IsKlineClosed    bool   `json:"x"`
	QuoteAssetVolume string `json:"q"`
	Buy              string `json:"V"`
	Quote            string `json:"Q"`
	Ignore           string `json:"B"`
}
