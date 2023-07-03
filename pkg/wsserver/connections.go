package wsserver

import (
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type Connections struct {
	conns map[string]*websocket.Conn
	mu    sync.Mutex
}

func connectionSlice() *Connections {
	return &Connections{
		conns: make(map[string]*websocket.Conn),
	}
}

func (conn *Connections) Add(userId string, ws *websocket.Conn) {
	conn.mu.Lock()
	conn.conns[userId] = ws
	conn.mu.Unlock()
}

func (conn *Connections) Remove(userId string) {
	conn.mu.Lock()
	conn.conns[userId].Close()
	delete(conn.conns, userId)
	conn.mu.Unlock()
}

func (conn *Connections) Broadcast(b []byte) {
	for _, ws := range conn.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println("warning: failed to broadcast message to websocket")
			}
		}(ws)
	}
}
