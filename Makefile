run:
	godotenv -f .env -- go run .


test:
	go test ./...