package share_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type ShareRepository interface {
	InsertOne(entity.ShareDB) (string, error)
	FindOne(filter, projection bson.M) (entity.ShareDB, error)
	Find(filter, projection bson.M) ([]entity.ShareDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.ShareDB, error)
	DeleteOne(filter bson.M) error
}
