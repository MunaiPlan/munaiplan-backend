.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans app

down:
	docker-compose down

rebuild:
	docker-compose down && docker-compose build && docker-compose up

swag:
	export PATH=$PATH:$HOME/go/bin
	swag init -g cmd/app/main.go