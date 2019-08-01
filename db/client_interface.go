package db

import "box-service/model"

/**
 *  := 	create date: 01-Jun-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type IRepository interface {
	Init(host string) error
	FindAvailable(capacity int32, maxWeight int32) (model.Box, error)
	Create(box model.Box) error
	Close()
}

var Client IRepository
