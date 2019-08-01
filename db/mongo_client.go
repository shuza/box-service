package db

import (
	"box-service/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	session *mgo.Session
}

func (repo *MongoRepository) Init(host string) error {
	session, err := mgo.Dial(host)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	repo.session = session

	return nil
}

func (repo *MongoRepository) FindAvailable(capacity int32, maxWeight int32) (model.Box, error) {
	filter := bson.M{
		"capacity":  bson.M{"$gte": capacity},
		"maxweight": bson.M{"$gte": maxWeight},
	}
	var box model.Box
	if err := repo.collection().Find(filter).One(&box); err != nil {
		return box, err
	}

	return box, nil
}

func (repo *MongoRepository) Create(box model.Box) error {
	return repo.collection().Insert(box)
}

func (repo *MongoRepository) collection() *mgo.Collection {
	return repo.session.DB("porter").C("boxes")
}

func (repo *MongoRepository) Close() {
	repo.session.Close()
}
