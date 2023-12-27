run: 
	go run cmd/server/main.go

setup:
	docker-compose build
	docker-compose up -d

stop:
	docker-compose down
