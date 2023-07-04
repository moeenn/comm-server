package server

import (
	"sync"

	"golang.org/x/net/websocket"
)

type Connections struct {
	Conns map[string]*websocket.Conn
	mu    sync.Mutex
}

func NewConnectionSlice() *Connections {
	return &Connections{
		Conns: make(map[string]*websocket.Conn),
	}
}

func (conn *Connections) Add(userId string, ws *websocket.Conn) {
	conn.mu.Lock()
	conn.Conns[userId] = ws
	conn.mu.Unlock()
}

func (conn *Connections) Remove(userId string) {
	conn.mu.Lock()
	conn.Conns[userId].Close()
	delete(conn.Conns, userId)
	conn.mu.Unlock()
}
