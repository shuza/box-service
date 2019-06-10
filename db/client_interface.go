package db

import (
	pb "github.com/shuza/box-service/proto"
)

/**
 *  := 	create date: 01-Jun-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type IRepository interface {
	Init(host string) error
	FindAvailable(spec *pb.Specification) (*pb.Box, error)
	Create(box *pb.Box) error
	Close()
}
