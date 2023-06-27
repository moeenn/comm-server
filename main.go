package main

import (
	"comm/config"
	"fmt"
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

	fmt.Printf("jwt: %v\ndb: %v\n", *config.JWTConfig, *config.DatabaseConfig)
}
