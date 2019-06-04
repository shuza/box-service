FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD box-service /app

CMD ["./box-service"]
