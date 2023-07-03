package main

import (
	"comm/config"
	"comm/pkg/wsserver"
	"comm/routes/notify"
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

	http.HandleFunc("/notify", notify.NotifyHandler(config.JWTConfig.ServerSecret))
	http.Handle("/ws", websocket.Handler(server.HandleWS(config.JWTConfig.ClientSecret)))

	log.Printf("info: starting server on port %s\n", config.ServerConfig.Port)
	http.ListenAndServe(config.ServerConfig.HostPort(), nil)
}
