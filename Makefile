run: 
	go run cmd/server/main.go

setup:
	go mod tidy
	docker-compose build
	docker-compose up -d

stop:
	docker-compose down
