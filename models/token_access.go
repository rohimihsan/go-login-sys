package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token_access struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	user_id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	token      string
	ip         string
	mac_addr   string
	created_at primitive.Timestamp
	updated_at primitive.Timestamp
}
