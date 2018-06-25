package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
}
