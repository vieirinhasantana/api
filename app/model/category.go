package model

import "gopkg.in/mgo.v2/bson"

type Category struct {
	ID        bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Descricao string        `json:"descricao"`
	Xyw       string        `json:"xyw"`
}
