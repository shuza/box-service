package main

import (
	"github.com/micro/go-micro"
	"github.com/shuza/box-service/db"
	pb "github.com/shuza/box-service/proto"
	"github.com/shuza/box-service/service"
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
	repo := &db.MongoRepository{}
	mongoHost := os.Getenv("MONGO_HOST")
	if err := repo.Init(mongoHost); err != nil {
		panic(err)
	}
	defer repo.Close()

	createDummyBox(repo)

	srv := micro.NewService(
		micro.Name("porter.box"),
		micro.Version("latest"),
	)
	srv.Init()

	//	Register our service with gRPC server
	//	this will tie our implementation into the auto-generated interface code
	//	for our protobuf edition
	boxService := service.NewBoxService(repo)
	pb.RegisterBoxServiceHandler(srv.Server(), &boxService)

}

func createDummyBox(repo db.IRepository) {
	repo.Create(&pb.Box{
		Id:        "box001",
		Name:      "First Box",
		MaxWeight: 200000,
		Capacity:  5000,
	})
}
