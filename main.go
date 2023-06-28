package main

import (
	"comm/config"
	"comm/pkg/wsserver"
	"fmt"
	"log"
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

	server := wsserver.New()

	http.HandleFunc("/test", middleware.ValidateToken(config.JWTConfig.Secret))
	http.Handle("/ws", websocket.Handler(server.HandleWS))

	log.Printf("info: starting websocket server on port %s\n", config.ServerConfig.Port)
	http.ListenAndServe(config.ServerConfig.HostPort(), nil)
}
