package common

import (
	"encoding/json"
	"github.com/workfoxes/kayo/internal/utils/ws"
	"github.com/workfoxes/kayo/pkg/log"
)

// BaseBroker stores the individual broker information
type BaseBroker struct {
	Name                 string
	IsWebSocketSupported bool
	Enabled              bool
	WS                   *ws.Conn
	Symbol               []string
	ItemChan             chan *Item
}

// OnWSConnected : Will be triggered when the websocket connection is opened for monitoring the trading api
func (b *BaseBroker) OnWSConnected(w *ws.Conn) {
	log.Info("Connected to " + b.Name + " Websocket Connection!")
}

func (b *BaseBroker) OnWSMessage(msg []byte, w *ws.Conn) {
	log.Info("New message new: " + string(msg))
}

func (b *BaseBroker) SendWSMessage(msg interface{}) {
	out, err := json.Marshal(msg)
	if err != nil {
		log.Error("SendWSMessage : " + err.Error())
	}
	err = b.WS.Send(ws.Msg{Body: out})
	if err != nil {
		log.Error("SendWSMessage : " + err.Error())
	}
}

func (b *BaseBroker) OnWSError(err error) {
	log.Error("Error from Websocket : ", err.Error())
}

func (b *BaseBroker) RegisterWebsocketClient(url string) {
	b.WS = &ws.Conn{
		OnConnected: b.OnWSConnected,
		OnMessage:   b.OnWSMessage,
		OnError:     b.OnWSError,
	}
	if err := b.WS.Dial(url, ""); err != nil {
		log.Error("Error While connecting to WebSocket in ", b.Name, " : ", err.Error())
	}
}

func (b *BaseBroker) Listen(symbol string, itemChan chan *Item) {
}

// Item : will have all the payload needed for the item with its config
type Item struct {
	Symbol       string
	Time         int64
	OpenPrice    string
	ClosePrice   string
	HighestPrice string
	LowestPrice  string
}
