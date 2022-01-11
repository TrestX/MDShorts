package db

import (
	"MdShorts/pkg/entity"
	search_repository "MdShorts/pkg/repository/search_repository"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	repo = search_repository.NewSearchRepository("search")
)

type searchService struct{}

func NewSearchService(repository search_repository.SearchRepository) SearchService {
	repo = repository
	return &searchService{}
}

func (*searchService) GetSearches(limit, skip int, userid string) ([]entity.SearchDB, error) {
	filter := bson.M{}
	if userid != "" {
		filter["user_id"] = userid
	}
	return repo.Find(filter, bson.M{})

}
