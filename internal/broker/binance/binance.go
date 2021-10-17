package binance

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/utils/ws"
	"github.com/workfoxes/kayo/pkg/log"
)

// Binance : is Object that will hold the connection and configuration with the Binance Service.
type Binance struct {
	common.BaseBroker
}

// Initialize : will initialize the broker config and setup
func (b *Binance) Initialize() {
	b.Name = common.Binance
	b.IsWebSocketSupported = true
}

// Listen : will listen the market data changes for the selected symbol
func (b *Binance) Listen(symbol string, itemChan chan *common.Item) {
	b.ItemChan = itemChan
	var symbols []string
	symbols = append(symbols, strings.Split(symbol, "|")...)
	log.Info(symbols)
	b.RegisterWebsocketClient(fmt.Sprintf("%s%s", StreamHostURL, RawStreamEndpoint))
	b.SendWSMessage(&WebSocketRequest{ID: 1000, Params: []string{"btcusdt@kline_1m"}, Method: "SUBSCRIBE"})
}

// OnWSMessage : will be triggered when a message is received from the Financial Broker
func (b *Binance) OnWSMessage(msg []byte, w *ws.Conn) {
	var binanceResponse WebSocketResponse
	binanceResponse.KPI.IsKlineClosed = false
	if err := json.Unmarshal(msg, &binanceResponse); err != nil {
		log.Error("Error from Websocket: " + err.Error())
	}
	if binanceResponse.KPI.IsKlineClosed {
		_item := convertToItem(&binanceResponse)
		log.Debug("New Trade Item: ", _item)
		b.ItemChan <- _item
	}
}

// convertToItem : convert the binance response to Kayo trading Item
func convertToItem(kline *WebSocketResponse) *common.Item {
	return &common.Item{
		Symbol:       kline.Symbol,
		Time:         kline.KPI.KStartTime,
		OpenPrice:    kline.KPI.OpenPrice,
		ClosePrice:   kline.KPI.ClosePrice,
		HighestPrice: kline.KPI.HighestPrice,
		LowestPrice:  kline.KPI.LowestPrice,
	}

}

func (b *Binance) OnWSError(err error) {
	log.Error("Error from Websocket: " + err.Error())
}

func (b *Binance) RegisterWebsocketClient(url string) {
	b.WS = &ws.Conn{
		OnConnected: b.OnWSConnected,
		OnMessage:   b.OnWSMessage,
		OnError:     b.OnWSError,
	}
	if err := b.WS.Dial(url, ""); err != nil {
		log.Error("Error While connecting to WebSocket in", b.Name, " :", err.Error())
	}
}
