package user_news_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type UserNewsCheckRepository interface {
	InsertOne(entity.UserNewsCheckDB) (string, error)
	FindOne(filter, projection bson.M) (entity.UserNewsCheckDB, error)
	Find(filter, projection bson.M) ([]entity.UserNewsCheckDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
