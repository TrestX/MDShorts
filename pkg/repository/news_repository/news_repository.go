package news_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type NewsRepository interface {
	InsertOne(entity.NewsDB) (string, error)
	FindOne(filter, projection bson.M) (entity.NewsDB, error)
	FindSort(filter, filter1, projection bson.M, limit, skip int) ([]entity.NewsDB, error)
	Find(filter, projection bson.M) ([]entity.NewsDB, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
	FindWithIDs(filter, projection bson.M) ([]entity.NewsDB, error)
}
