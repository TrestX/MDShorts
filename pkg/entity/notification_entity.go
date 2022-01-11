package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageData struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Title      string             `bson:"title" json:"title"`
	Body       string             `bson:"body" json:"body"`
	Topic      string             `bson:"topic" json:"topic"`
	UserId     string             `bson:"userId" json:"userId"`
	CategoryId string             `bson:"categoryId" json:"categoryId"`
	SentTime   time.Time          `bson:"sentTime" json:"sentTime"`
	Status     string             `bson:"status" json:"status"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
