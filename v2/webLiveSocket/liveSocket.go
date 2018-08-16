package webLiveSocket

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WebLiveSocket struct {
	longCon      *websocket.Conn
	transferData chan []byte
	closeChan    chan byte
	isClosed     bool
	mutex        sync.Mutex
}

func InitMySocket(conn *websocket.Conn) (sk *WebLiveSocket) {
	sk = &WebLiveSocket{
		longCon:      conn,
		transferData: make(chan []byte, 1024),
		closeChan:    make(chan byte, 1),
		isClosed:     false,
	}
	return
}

func (c *WebLiveSocket) close() (err error) {
	if err = c.longCon.Close(); err != nil {
		return
	}
	c.mutex.Lock()
	if !c.isClosed {
		close(c.closeChan)
		c.isClosed = true
	}
	c.mutex.Unlock()
	return
}

func (c *WebLiveSocket) ReadLoop() {
	for {
		var (
			err  error
			data []byte
		)
		if _, data, err = c.longCon.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case c.transferData <- data:
		case <-c.closeChan:
			goto ERR
		}
	}
ERR:
	c.close()
}

func (c *WebLiveSocket) WirteLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-c.transferData:
			if err = c.longCon.WriteMessage(websocket.TextMessage, data); err != nil {
				goto ERR
			}
		case <-c.closeChan:
			goto ERR
		}
	}
ERR:
	c.close()
}
