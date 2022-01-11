package unregistered_user_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type UnRegisteredUserRepository interface {
	InsertOne(entity.UnRegisteredUsersDB) (string, error)
	FindOne(filter, projection bson.M) (entity.UnRegisteredUsersDB, error)
	Find(filter, projection bson.M) ([]entity.UnRegisteredUsersDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.UnRegisteredUsersDB, error)
	DeleteOne(filter bson.M) error
}
