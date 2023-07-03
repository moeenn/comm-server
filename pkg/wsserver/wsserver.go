package wsserver

import (
	"comm/pkg/error"
	"comm/pkg/services/auth"
	"encoding/json"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const BUF_SIZE = 2048

type WSServer struct {
	conns *Connections
}

func New() *WSServer {
	return &WSServer{
		conns: connectionSlice(),
	}
}

func (s *WSServer) HandleWS(jwtSecret string) websocket.Handler {
	return func(ws *websocket.Conn) {
		userId, err := auth.ValidateQueryToken(jwtSecret, ws.Request())
		if err != nil {
			log.Printf("warning: tried to connect websocket without providing valid auth token")

			errorResponse := error.ErrorResponse{
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
	defer s.conns.Remove(userId)

	for {
		n, err := ws.Read(buf)
		if err != nil {

			// handle situation where the client closes the connection
			if err == io.EOF {
				log.Println("info: client disconnected socket")

				// breaking out of the loop will close the websocket
				break
			}

			// if message read fails for any reason, log the error but don't
			// disconnect the socket
			log.Printf("error: %s\n", err.Error())
			continue
		}

		// TODO: process the incoming message
		msgBytes := buf[:n]
		s.conns.Broadcast(msgBytes)
	}
}
