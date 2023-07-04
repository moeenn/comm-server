run:
	godotenv -f .env -- go run .


test:
	go test ./...


clientgen:
	sqlc generate