package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User_profile struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	address    string
	city       string
	province   string
	country    string
	phone      string
	created_at time.Time
	updated_at time.Time
}
