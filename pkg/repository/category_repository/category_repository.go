package category_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type CategoryRepository interface {
	InsertOne(entity.CategoryDB) (string, error)
	FindOne(filter, projection bson.M) (entity.CategoryDB, error)
	Find(filter, projection bson.M) ([]entity.CategoryDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	FindWithIDs(filter, projection bson.M) ([]entity.CategoryDB, error)
	DeleteOne(filter bson.M) error
}
