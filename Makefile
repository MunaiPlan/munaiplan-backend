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

run-model:
	docker pull kabdulaset/munai-models && docker run -d -p 8001:80 kabdulaset/munai-models

migrate-diff:
	atlas migrate diff --env gorm

