//+build !test

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/workfoxes/kayo/internal/utils/ws"
	"github.com/workfoxes/kayo/pkg/log"
)

func main() {

	conn := ws.Conn{
		// Fires when the connection is established
		OnConnected: func(w *ws.Conn) {
			fmt.Println("Connected!")
		},
		// Fires when a new message arrives from the server
		OnMessage: func(msg []byte, w *ws.Conn) {
			fmt.Printf("New message: %s\n", msg)
		},
		// Fires when an error occurs and connection is closed
		OnError: func(err error) {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		},
		// Ping interval in secs (optional)
		//PingIntervalSecs: 0,
		// Ping message to send (optional)
		//PingMsg: []byte("PING"),
	}

	err := conn.Dial("wss://stream.binance.com:9443/ws/!ticker@arr", "")
	time.Sleep(time.Millisecond * 10000000)
	if err != nil {
		log.Fatal(err)
	}
}
