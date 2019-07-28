package main

import (
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/shuza/box-service/db"
	pb "github.com/shuza/box-service/proto"
	"github.com/shuza/box-service/service"
	log "github.com/sirupsen/logrus"
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

	registry := consul.NewRegistry()
	srv := k8s.NewService(
		micro.Name("porter.box"),
		micro.Version("latest"),
		micro.Registry(registry),
	)
	srv.Init()

	//	Register our service with gRPC server
	//	this will tie our implementation into the auto-generated interface code
	//	for our protobuf edition
	boxService := service.NewBoxService(repo)
	pb.RegisterBoxServiceHandler(srv.Server(), &boxService)

	if err := srv.Run(); err != nil {
		log.Warnf("srv Run  Error  :  %v\n", err)
	}
}

func createDummyBox(repo db.IRepository) {
	repo.Create(&pb.Box{
		Id:        "box001",
		Name:      "First Box",
		MaxWeight: 200000,
		Capacity:  5000,
	})
}
