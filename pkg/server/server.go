package server

import (
	"io"
	"log"

	"comm/config"
	"comm/pkg/error"
	mw "comm/pkg/middleware"
	"comm/pkg/services/auth"
	"comm/routes/notify"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/websocket"
)

type Server struct {
	config      *config.Config
	Router      *chi.Mux
	connections *Connections
}

func New(config *config.Config) *Server {
	server := &Server{
		config:      config,
		Router:      chi.NewRouter(),
		connections: NewConnectionSlice(),
	}

	// register all middleware here
	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(mw.ValidateBearerToken(server.config.JWTConfig.ServerSecret))
	server.Router.Handle("/ws", websocket.Handler(server.WSHandler))

	// register all routes here
	server.Router.Post("/api/notify", notify.NotifyHandler)

	return server
}

// send a message to all connected websockets
func (server *Server) Broadcast(b []byte) {
	for _, conn := range server.connections.Conns {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(b); err != nil {
				log.Println("warning: failed to broadcast message to websocket")
			}
		}(conn)
	}
}

func (server *Server) connectionReadLoop(conn *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("connection closed by client")
				break
			}

			log.Printf("data read error: %s\n", err.Error())
			continue
		}

		// TODO: do something with the incoming messages
		msgBytes := buf[:n]
		go func(conn *websocket.Conn, payload []byte) {
			if _, err := conn.Write(payload); err != nil {
				log.Printf("conn write failed: %s\n", err.Error())
			}
		}(conn, msgBytes)
	}
}

func (server *Server) WSHandler(conn *websocket.Conn) {
	// authenticate the request
	userId, err := auth.ValidateQueryToken(server.config.JWTConfig.ClientSecret, conn.Request())
	if err != nil {
		errorResponse := error.ErrorResponse{
			Error: err.Error(),
		}

		resJson, _ := json.Marshal(errorResponse)
		conn.Write(resJson)

		return
	}

	server.connections.Add(userId, conn)
	defer server.connections.Remove(userId)

	// activate the socket
	server.connectionReadLoop(conn)
}
