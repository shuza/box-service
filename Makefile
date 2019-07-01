include .env

protobuf:
	protoc -I. --go_out=plugins=micro:. ./proto/*.proto

.PHONY: build docker image
docker_build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/box-service:$(APP_VERSION) .

.PHONY: run docker image
docker_run: docker_build
	docker run -p 8081:8081 \
	-e MICRO_SERVER_ADDRESS=:8081 \
	-e MONGO_HOST=$(MONGO_HOST) \
	shuzasa/box-service:$(APP_VERSION)

.PHONY: publish in DockerHub
docker_push: docker_build
	docker push shuzasa/box-service:$(APP_VERSION)

run:
	MICRO_SERVER_ADDRESS=$(MICRO_SERVER_ADDRESS) \
	MONGO_HOST=$(MONGO_HOST) \
	go run main.go