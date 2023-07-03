package main

import (
	"comm/config"
	"comm/pkg/wsserver"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	server := wsserver.New()
	http.Handle("/ws", websocket.Handler(server.HandleWS(config.JWTConfig.Secret)))

	log.Printf("info: starting websocket server on port %s\n", config.ServerConfig.Port)
	http.ListenAndServe(config.ServerConfig.HostPort(), nil)
}
