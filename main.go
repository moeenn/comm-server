package main

import (
	"comm/config"
	"comm/pkg/logger"
	"comm/pkg/wsserver"
	"fmt"
	"net/http"

	"comm/pkg/middleware"

	"golang.org/x/net/websocket"
)

const (
	TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIzMzI2YTNlZS0zMGFlLTRjMGQtYjBkYS1mYzBjZjhjM2RjYTEiLCJzY29wZSI6IkFVVEgiLCJ1c2VyUm9sZSI6IlVTRVIiLCJpYXQiOjE2ODc4NDM0NjN9.rLqMp1NpeHYg4iTcLKk354tje028sx-3PFUojZfW698"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	log := logger.New()
	server := wsserver.New(log)

	http.HandleFunc("/test", middleware.ValidateToken(config.JWTConfig.Secret))
	http.Handle("/ws", websocket.Handler(server.HandleWS))

	log.Info("starting websocket server on port " + config.ServerConfig.Port)
	http.ListenAndServe(config.ServerConfig.HostPort(), nil)
}
