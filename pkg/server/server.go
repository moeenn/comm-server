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

	"comm/database"

	"comm/pkg/notification"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/websocket"
)

type Server struct {
	config              *config.Config
	db                  *database.Database
	Router              *chi.Mux
	connections         *Connections
	notificationChannel chan notification.Notification
}

func New(config *config.Config, db *database.Database) *Server {
	server := &Server{
		config:              config,
		db:                  db,
		Router:              chi.NewRouter(),
		connections:         NewConnectionSlice(),
		notificationChannel: make(chan notification.Notification, 10),
	}

	// register all middleware here
	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(mw.ValidateBearerToken(server.config.JWTConfig.ServerSecret))
	server.Router.Handle("/ws", websocket.Handler(server.WSHandler))

	// register all routes here
	server.Router.Post("/api/notify", notify.NotifyHandler(server.notificationChannel))

	go server.processNotifications()
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

		// TODO: check if channel has any pending notification and push them out

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

func (s *Server) processNotifications() {
	for notification := range s.notificationChannel {
		for _, userId := range notification.UserIds {
			conn, ok := s.connections.Conns[userId]
			if !ok {
				// user is not online

				continue
			}

			conn.Write(notification.Payload)
		}
	}
}
