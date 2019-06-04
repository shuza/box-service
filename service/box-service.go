package service

import (
	"context"
	"github.com/shuza/box-service/db"
	pb "github.com/shuza/box-service/proto"
	log "github.com/sirupsen/logrus"
)

/**
 *  := 	create date: 01-Jun-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type boxService struct {
	repo db.IRepository
}

func NewBoxService(repo db.IRepository) boxService {
	service := boxService{repo: repo}
	return service
}

func (s *boxService) FindAvailableBox(ctx context.Context, req *pb.Specification) (*pb.Response, error) {
	box, err := s.repo.FindAvailable(req)
	if err != nil {
		log.Warnf("repo Find available Error :  %v", err)
		return nil, err
	}

	return &pb.Response{Box: box}, nil
}
