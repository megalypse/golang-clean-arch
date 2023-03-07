run-server:
	docker-compose up -d
	go run cmd/server/main.go
