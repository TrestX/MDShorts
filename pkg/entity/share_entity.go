package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShareDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"shareId"`
	UserId     string             `bson:"user_id" json:"userId,omitempty"`
	NewsId     string             `bson:"newsId" json:"newsId"`
	SharedVia  string             `bson:"shared_via" json:"sharedVia"`
	SharedTime time.Time          `bson:"shared_time" json:"sharedTime"`
}
