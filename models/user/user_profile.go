package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User_profile struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	user_id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	dynamic_field string
	value         string
	created_at    primitive.Timestamp
	updated_at    primitive.Timestamp
}
