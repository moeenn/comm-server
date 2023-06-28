package wsserver

import (
	connections "comm/pkg/wsserver/connections"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const BUF_SIZE = 2048

type WSServer struct {
	conns *connections.Connections
}

func New() *WSServer {
	return &WSServer{
		conns: connections.New(),
	}
}

func (s *WSServer) HandleWS(ws *websocket.Conn) {
	log.Printf("info: new connection: %s", ws.RemoteAddr())
	s.conns.Add(ws)
	s.readLoop(ws)
}

func (s *WSServer) readLoop(ws *websocket.Conn) {
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := ws.Read(buf)
		if err != nil {

			/* handle situation where the client closes the connection */
			if err == io.EOF {
				log.Println("info: client disconnected socket")
				break
			}

			/* TODO: do proper error handling */
			log.Printf("error: %s\n", err.Error())

			/* breaking out of the loop will close the websocket */
			continue
		}

		/* TODO: process the incoming message */
		msgBytes := buf[:n]
		s.Broadcast(msgBytes)
	}

	/* delete socket when it closes */
	s.conns.Remove(ws)
}

func (s *WSServer) Broadcast(b []byte) {
	s.conns.BroadCast(b)
}
