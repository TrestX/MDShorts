package user_news_update_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type UserNewsUpdateRepository interface {
	InsertOne(entity.UserNewsUpdateDB) (string, error)
	FindOne(filter, projection bson.M) (entity.UserNewsUpdateDB, error)
	Find(filter, projection bson.M) ([]entity.UserNewsUpdateDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
