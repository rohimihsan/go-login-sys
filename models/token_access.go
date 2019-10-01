package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Token_access struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	user_id    primitive.ObjectID `json:"user_id"`
	token      string
	ip         string
	mac_addr   string
	created_at time.Time
	updated_at time.Time
}
