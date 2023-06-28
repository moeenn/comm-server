package wsserver

import (
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type Connections struct {
	conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func New() *Connections {
	return &Connections{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (conn *Connections) Add(ws *websocket.Conn) {
	conn.mu.Lock()
	conn.conns[ws] = true
	conn.mu.Unlock()
}

func (conn *Connections) Remove(ws *websocket.Conn) {
	conn.mu.Lock()
	delete(conn.conns, ws)
	conn.mu.Unlock()
}

func (conn *Connections) BroadCast(b []byte) {
	for ws := range conn.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println("warning: failed to broadcast message to websocket")
			}
		}(ws)
	}
}
