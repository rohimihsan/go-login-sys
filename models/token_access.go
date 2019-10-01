package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Token_access struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User_id    primitive.ObjectID `json:"user_id"`
	Token      string             `json:"token"`
	Ip         string             `json:"ip"`
	Mac_addr   string             `json:"mac_addr"`
	Valid      bool               `json:"valid"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
