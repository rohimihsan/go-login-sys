package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	firstname   string
	lastname    string
	username    string
	email       string
	user_number string
	created_at  primitive.Timestamp
	updated_at  primitive.Timestamp
}
