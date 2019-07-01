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

func (s *boxService) FindAvailableBox(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	box, err := s.repo.FindAvailable(req)
	if err != nil {
		log.Warnf("repo Find available Error :  %v", err)
		return err
	}
	resp.Box = box
	return nil
}

func (s *boxService) Create(ctx context.Context, req *pb.Box, resp *pb.Response) error {
	log.Infof("Create service called")
	if err := s.repo.Create(req); err != nil {
		log.Warnf("repo create box Error  :   %v", err)
		return err
	}
	resp.Created = true
	resp.Box = req
	return nil
}
