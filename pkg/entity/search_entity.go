package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"shareId"`
	UserId     string             `bson:"user_id" json:"userId,omitempty"`
	Search     string             `bson:"search" json:"search"`
	SearchedAt time.Time          `bson:"searched_at" json:"searchedAt"`
}
