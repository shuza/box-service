package db

import (
	"errors"
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
	FindAvailable(spec *pb.Specification) (*pb.Box, error)
	Create(box *pb.Box) error
}

type Repository struct {
	boxes []*pb.Box
}

func (repo *Repository) FindAvailable(spec *pb.Specification) (*pb.Box, error) {
	for _, box := range repo.boxes {
		if spec.Capacity <= box.Capacity && spec.MaxWeight <= box.MaxWeight {
			return box, nil
		}
	}

	return nil, errors.New("No box found by that spec")
}

func (repo *Repository) Create(box *pb.Box) error {
	repo.boxes = append(repo.boxes, box)
	return nil
}
