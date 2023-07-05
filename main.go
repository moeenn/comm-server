package main

import (
	"comm/config"
	"comm/database"
	"comm/pkg/server"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to load application config: %s\n", err.Error())
		os.Exit(1)
	}

	db, err := database.Connect(conf.DatabaseConfig.URI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: db connection failed: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	server := server.New(conf, db)
	if err := http.ListenAndServe(conf.ServerConfig.HostPort(), server.Router); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
