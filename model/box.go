package model

import "gopkg.in/mgo.v2/bson"

type Box struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Capacity  int32         `json:"capacity"`
	MaxWeight int32         `json:"max_weight"`
	Name      string        `json:"name"`
	Available bool          `json:"available"`
	OwnerId   string        `json:"owner_id"`
}
