package wsserver

import (
	"comm/pkg/logger"
	connections "comm/pkg/wsserver/connections"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

const BUF_SIZE = 2048

type WSServer struct {
	logger *logger.Logger
	conns  *connections.Connections
}

func New(logger *logger.Logger) *WSServer {
	return &WSServer{
		logger: logger,
		conns:  connections.New(),
	}
}

func (s *WSServer) HandleWS(ws *websocket.Conn) {
	s.logger.Info(fmt.Sprintf("new incoming connection: %s", ws.RemoteAddr()))
	s.conns.Add(ws)
	s.readLoop(ws)
}

func (s *WSServer) readLoop(ws *websocket.Conn) {
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := ws.Read(buf)
		if err != nil {

			/** handle situation where the client closes the connection */
			if err == io.EOF {
				s.logger.Info("client disconnected socket")
				break
			}

			/** TODO: do proper error handling */
			s.logger.Error(fmt.Sprintf("read error: %v+", err))

			/** breaking out of the loop will close the websocket */
			continue
		}

		/** echo the incoming message to all connected websockets */
		msgBytes := buf[:n]
		s.Broadcast(msgBytes)
	}

	/* delete socket when it closes */
	s.conns.Remove(ws)
}

func (s *WSServer) Broadcast(b []byte) {
	s.conns.BroadCast(b, s.logger)
}
