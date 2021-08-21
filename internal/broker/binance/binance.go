package binance

import (
	"encoding/json"
	"fmt"
	"github.com/workfoxes/kayo/internal/broker/default"
	"github.com/workfoxes/kayo/internal/utils/ws"
	"github.com/workfoxes/kayo/pkg/log"
)

var obj WebSocketResponse

type Binance struct {
	_default.BaseBroker
}

func (b *Binance) Initialize() {
	b.Name = _default.Binance
	b.IsWebSocketSupported = true
}

func (b *Binance) Listen(itemChan chan *_default.Item) {
	b.ItemChan = itemChan
	b.RegisterWebsocketClient(fmt.Sprintf("%s%s", StreamHostURL, RawStreamEndpoint))
	b.SendWSMessage(&WebSocketRequest{ID: 1000, Params: []string{"btcusdt@kline_1m"}, Method: "SUBSCRIBE"})
}

func (b *Binance) OnWSMessage(msg []byte, w *ws.Conn) {
	obj.KPI.IsKlineClosed = false
	if err := json.Unmarshal(msg, &obj); err != nil {
		log.Error("Error from Websocket: " + err.Error())
	}
	if obj.KPI.IsKlineClosed {
		log.Info("New message: ", obj)
		_item := convertToItem(obj)
		b.ItemChan <- _item
	}
}

func convertToItem(kline WebSocketResponse) *_default.Item {
	return &_default.Item{
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
	err := b.WS.Dial(url, "")
	if err != nil {
		log.Error("Error While connecting to WebSocket in %s : %s ", b.Name, err.Error())
	}
}
