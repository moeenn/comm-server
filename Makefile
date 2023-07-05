run:
	godotenv -f .env -- go run .


check:
	sqlc compile && go vet ./...


test:
	go test ./...


clientgen:
	sqlc generate