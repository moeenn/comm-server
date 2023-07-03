package wsserver

import (
	"comm/pkg/service"
	connections "comm/pkg/wsserver/connections"
	"encoding/json"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const BUF_SIZE = 2048

type ErrorResponse struct {
	Error string `json:"error"`
}

type WSServer struct {
	conns *connections.Connections
}

func New() *WSServer {
	return &WSServer{
		conns: connections.New(),
	}
}

func (s *WSServer) HandleWS(jwtSecret string) websocket.Handler {
	return func(ws *websocket.Conn) {
		userId, err := service.ValidateToken(jwtSecret, ws.Request())
		if err != nil {
			log.Printf("warning: tried to connect websocket without providing valid auth token")

			errorResponse := ErrorResponse{
				Error: "error: " + err.Error(),
			}

			resJson, _ := json.Marshal(errorResponse)
			ws.Write(resJson)

			return
		}

		log.Printf("info: new connection: %s", ws.RemoteAddr())
		s.conns.Add(userId, ws)
		s.readLoop(userId, ws)
	}
}

func (s *WSServer) readLoop(userId string, ws *websocket.Conn) {
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := ws.Read(buf)
		if err != nil {

			// handle situation where the client closes the connection
			if err == io.EOF {
				log.Println("info: client disconnected socket")
				break
			}

			// if message read fails for any reason, log the error but don't
			// disconnect the socket
			log.Printf("error: %s\n", err.Error())

			// breaking out of the loop will close the websocket
			continue
		}

		/* TODO: process the incoming message */
		msgBytes := buf[:n]
		s.Broadcast(msgBytes)
	}

	// delete socket when it closes
	s.conns.Remove(userId)
}

func (s *WSServer) Broadcast(b []byte) {
	s.conns.BroadCast(b)
}
