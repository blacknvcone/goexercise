package models

import "gopkg.in/mgo.v2/bson"

type Base struct {
	Id        bson.ObjectId `bson:"_id"`
	CreatedAt string        `bson:"created_at"`
}
