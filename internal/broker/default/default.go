package _default

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
	//Verbose                       bool
	//LoadedByConfig                bool
	//SkipAuthCheck                 bool
	//HTTPTimeout                   time.Duration
	//HTTPUserAgent                 string
	//HTTPRecording                 bool
	//HTTPDebugging                 bool
	//WebsocketResponseCheckTimeout time.Duration
	//WebsocketResponseMaxLimit     time.Duration
	//WebsocketOrderBookBufferLimit int64
	//SettingsMutex                 sync.RWMutex
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
	log.Error("Error from Websocket: %s\n", err.Error())
}

func (b *BaseBroker) RegisterWebsocketClient(url string) {
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

// Item : will have all the payload needed for the item with its config
type Item struct {
	Symbol       string
	Time         int64
	OpenPrice    string
	ClosePrice   string
	HighestPrice string
	LowestPrice  string
}

type TradeConfig struct {
}
