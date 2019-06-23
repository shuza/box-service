include .env

protobuf:
	protoc -I . --go_out=plugins=grpc:. ./proto/*.proto

.PHONY: build docker image
docker_build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/box-service .

.PHONY: run docker image
docker_run: docker_build
	docker run -p 8081:8081 -e PORT=:8081 -e MONGO_HOST=localhost:27017 shuzasa/box-service

run:
	PORT=$(PORT) \
	MONGO_HOST=$(MONGO_HOST) \
	go run main.go