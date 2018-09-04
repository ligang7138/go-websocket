package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket/v2/webLiveSocket"
	"net/http"
)

var (
	// http升级wobsocket的握手交互
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ms   *webLiveSocket.WebLiveSocket
		conn *websocket.Conn
		err  error
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		fmt.Println("升级websocket协议失败")
		conn.Close()
		return
	}
	ms = webLiveSocket.InitMySocket(conn)
	go ms.ReadLoop()
	go ms.WirteLoop()
}

func main() {

	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
