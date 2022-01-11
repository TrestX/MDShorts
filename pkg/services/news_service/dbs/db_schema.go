package db

import (
	"MdShorts/pkg/entity"
)

type NewsService interface {
	GetNews(userId string, limit, skip int) ([]entity.NewsDB, error)
	GetNewsByID(newsID, userId string, limit, skip int) ([]entity.NewsDB, error)
	GetGlobalNews(country, language string, limit, skip int) ([]entity.NewsDB, error)
	GetSearchNews(search, uid string, limit, skip int) ([]entity.NewsDB, error)
	GetFeaturedNews(limit, skip int) ([]entity.NewsDB, error)
	GetTrendingNews(limit, skip int) ([]entity.NewsDB, error)
	GetAllNews(limit, skip int) ([]entity.NewsDB, error)
}
