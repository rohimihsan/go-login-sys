package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dynamic_field struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	name       string
	desc       string
	data_type  string
	field_type int //1: user profile
	created_at primitive.Timestamp
	updated_at primitive.Timestamp
}
