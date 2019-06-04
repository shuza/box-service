protobuf:
	protoc -I . --go_out=plugins=grpc:. ./proto/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/box-service .

run:
	docker run -p 8081:8081 -e PORT=:8081 shuzasa/box-service
