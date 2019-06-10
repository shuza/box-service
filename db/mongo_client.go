package db

import (
	pb "github.com/shuza/box-service/proto"
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

func (repo *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Box, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	var box *pb.Box
	if err := repo.collection().Find(filter).One(&box); err != nil {
		return nil, err
	}

	return box, nil
}

func (repo *MongoRepository) Create(box *pb.Box) error {
	return repo.collection().Insert(box)
}

func (repo *MongoRepository) collection() *mgo.Collection {
	return repo.session.DB("porter").C("boxes")
}

func (repo *MongoRepository) Close() {
	repo.session.Close()
}
