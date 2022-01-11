package notification_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type NotificationRepository interface {
	InsertOne(entity.MessageData) (string, error)
	FindOne(filter, projection bson.M) (entity.MessageData, error)
	Find(filter, projection bson.M) ([]entity.MessageData, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	FindSort(filter, filter1, projection bson.M, limit, skip int) ([]entity.MessageData, error)
}
