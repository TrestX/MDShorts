package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeeFewerStoriesDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"fewerStoriesId"`
	UserId     string             `bson:"userId" json:"userId"`
	SourceName string             `bson:"sourceName" json:"sourceName"`
	AddedTime  time.Time          `bson:"added_time" json:"addedTime"`
}
