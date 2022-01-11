package db

import (
	"MdShorts/pkg/entity"
)

type SearchService interface {
	GetSearches(limit, skip int, userid string) ([]entity.SearchDB, error)
}
