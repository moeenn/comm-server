package main

import (
	"comm/config"
	"comm/pkg/server"
	"fmt"
	"net/http"
	"os"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to load application config: %s\n", err.Error())
		os.Exit(1)
	}

	server := server.New(conf)
	http.ListenAndServe(conf.ServerConfig.HostPort(), server.Router)
}
