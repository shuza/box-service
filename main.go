package main

import (
	"github.com/shuza/box-service/db"
	pb "github.com/shuza/box-service/proto"
	"github.com/shuza/box-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"os"
)

/**
 *  := 	create date: 01-Jun-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

func main() {
	repo := &db.Repository{}
	port := os.Getenv("PORT")

	//	Create initial box
	repo.Create(&pb.Box{
		Id:"box001",
		Name:"First Box",
		MaxWeight:200000,
		Capacity:5000,
	})

	//	setup gRPC server
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen  :  ", err)
	}

	s := grpc.NewServer()

	//	Register our service with gRPC server
	//	this will tie our implementation into the auto-generated interface code
	//	for our protobuf edition
	boxService := service.NewBoxService(repo)
	pb.RegisterBoxServiceServer(s, &boxService)

	reflection.Register(s)

	log.Println("Running on port :  ", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to server  :  ", err)
	}
}
