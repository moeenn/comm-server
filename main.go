package main

import (
	"comm/config"
	"comm/database"
	"comm/pkg/server"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error: failed to load application config: %s\n", err.Error())
	}

	db, err := database.Connect(conf.DatabaseConfig.URI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	server := server.New(conf, db)
	if err := http.ListenAndServe(conf.ServerConfig.HostPort(), server.Router); err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
