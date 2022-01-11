package user_login_history_repository

import (
	"go.mongodb.org/mongo-driver/bson"
)

type UserLoginHistoryRepository interface {
	InsertOne(interface{}) (string, error)
	FindOne(filter, projection bson.M) (interface{}, error)
	Find(filter, projection bson.M) ([]interface{}, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
