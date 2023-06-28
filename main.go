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
