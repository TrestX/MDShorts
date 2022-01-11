package search_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type SearchRepository interface {
	InsertOne(entity.SearchDB) (string, error)
	FindOne(filter, projection bson.M) (entity.SearchDB, error)
	Find(filter, projection bson.M) ([]entity.SearchDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}
