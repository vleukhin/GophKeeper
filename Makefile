build:
	go build -o bin/gophkeeper-client cmd/client/main.go && chmod +x ./bin/gophkeeper-client

up:
	docker compose up -d