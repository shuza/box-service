package main

import (
	"box-service/api"
	"box-service/db"
	"fmt"
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
	initDb()
	defer db.Client.Close()

	r := api.NewGinEngine()
	fmt.Println("Box service is running on port 8081 ....")
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

func initDb() {
	db.Client = &db.MongoRepository{}
	if err := db.Client.Init(os.Getenv("MONGO_HOST")); err != nil {
		panic(err)
	}
}
