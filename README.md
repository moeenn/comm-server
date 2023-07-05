# comm-server
A web-sockets server written in Golang.


## Preparation
```bash
# for loading env variables from a file
$ go install github.com/joho/godotenv/cmd/godotenv@latest

# for generating database client from raw sql queries
$ go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# generate code for database client
$ make clientgen
```

## Usage
```bash
$ make run
```