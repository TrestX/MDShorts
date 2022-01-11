package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryDB struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Status       string             `bson:"status,omitempty" json:"status"`
	CategoryName string             `bson:"category_name" json:"categoryName"`
	PreSignedUrl string             `bson:"pre_signed_url" json:"imageUrl"`
	AddedTime    time.Time          `bson:"added_time" json:"addedTime"`
	UpdatedTime  time.Time          `bson:"updated_time" json:"updatedTime"`
}
