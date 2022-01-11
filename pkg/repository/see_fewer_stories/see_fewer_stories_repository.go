package share_repository

import (
	"MdShorts/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type SeeFewerStoriesRepository interface {
	InsertOne(entity.SeeFewerStoriesDB) (string, error)
	FindOne(filter, projection bson.M) (entity.SeeFewerStoriesDB, error)
	Find(filter, projection bson.M) ([]entity.SeeFewerStoriesDB, error)
	UpdateOne(filter, update bson.M) (string, error)
}
