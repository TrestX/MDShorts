package db

import (
	"MdShorts/pkg/entity"
	news_repository "MdShorts/pkg/repository/news_repository"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	category_repository "MdShorts/pkg/repository/category_repository"
	catdb "MdShorts/pkg/services/category_service/dbs"

	search_repository "MdShorts/pkg/repository/search_repository"
	see_fewer_stories "MdShorts/pkg/repository/see_fewer_stories"
)

var (
	repo = news_repository.NewNewsRepository("news")
)
var (
	categoryService = catdb.NewCategoryService(category_repository.NewCategoryRepository("category"))
)
var (
	search_repo = search_repository.NewSearchRepository("search")
)
var (
	seefewer_repo = see_fewer_stories.NewSeeFewerStoriesRepository("see_fewer_stories")
)

type newsService struct{}

func NewNewsService(repository news_repository.NewsRepository) NewsService {
	repo = repository
	return &newsService{}
}

func (*newsService) GetGlobalNews(language, country string, limit, skip int) ([]entity.NewsDB, error) {
	news, err := repo.FindSort(bson.M{"category": bson.M{"$in": []string{"global", "international"}}, "publishedAt": bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -1))}}, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
	if err != nil || len(news) < 1 || news == nil {
		news, _ = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	return news, nil
}

//
func (*newsService) GetSearchNews(search, uid string, limit, skip int) ([]entity.NewsDB, error) {
	splitQuery := strings.Split(search, " ")
	filter := bson.M{}
	var searc entity.SearchDB
	searc.ID = primitive.NewObjectID()
	searc.SearchedAt = time.Now()
	searc.Search = search
	searc.UserId = uid
	_, _ = search_repo.InsertOne(searc)
	var searchList bson.A
	for i := 0; i < len(splitQuery); i++ {
		searchList = append(searchList, bson.M{"title": bson.M{"$regex": splitQuery[i], "$options": "i"}})
		searchList = append(searchList, bson.M{"description": bson.M{"$regex": splitQuery[i], "$options": "i"}})
		searchList = append(searchList, bson.M{"sourceName": bson.M{"$regex": splitQuery[i], "$options": "i"}})
		searchList = append(searchList, bson.M{"author": bson.M{"$regex": splitQuery[i], "$options": "i"}})
	}
	filter["$or"] = searchList
	return repo.FindSort(filter, bson.M{"_id": -1}, bson.M{}, limit, skip)
}
func (*newsService) GetFeaturedNews(limit, skip int) ([]entity.NewsDB, error) {
	filter := bson.M{}
	filter["featured"] = true
	filter["publishedAt"] = bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now().Add(-12))}
	news, err := repo.FindSort(filter, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
	if err != nil || len(news) < 1 || news == nil {
		news, _ = repo.FindSort(bson.M{"publishedAt": bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -1))}}, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	return news, nil
}

func (*newsService) GetTrendingNews(limit, skip int) ([]entity.NewsDB, error) {
	filter := bson.M{}
	filter["trending"] = true
	filter["publishedAt"] = bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now().Add(-8))}
	news, err := repo.FindSort(filter, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
	if err != nil || len(news) < 1 || news == nil {
		news, _ = repo.FindSort(bson.M{"publishedAt": bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -1))}}, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	return news, nil
}

func (*newsService) GetNews(userId string, limit, skip int) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	var err error
	if userId == "" {
		news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
		if err != nil {
			return []entity.NewsDB{}, err
		}
	} else {
		catdata, err := categoryService.GetCategoryForUser(userId)
		if err != nil {
			news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
			if err != nil {
				return []entity.NewsDB{}, err
			}
		} else {
			l := []string{}
			for _, cat := range catdata {
				l = append(l, cat.CategoryName)
			}
			seefewer, _ := seefewer_repo.Find(bson.M{"userId": userId}, bson.M{})
			sfs := []string{}
			if len(seefewer) > 0 {
				for i := 0; i < len(seefewer); i++ {
					sfs = append(sfs, seefewer[i].SourceName)
				}
			}
			filter := bson.M{}
			filter["category"] = bson.M{"$in": l}
			filter["sourceName"] = bson.M{"$nin": sfs}
			news, err = repo.FindSort(filter, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
			if err != nil || len(news) == 0 {
				news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
				if err != nil {
					return []entity.NewsDB{}, err
				}
			}

		}
	}
	if err != nil {
		return []entity.NewsDB{}, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	return news, nil
}

func (*newsService) GetAllNews(limit, skip int) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	var err error
	news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
	if err != nil {
		return []entity.NewsDB{}, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	return news, nil
}

func (*newsService) GetNewsByID(newsID, userId string, limit, skip int) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	var err error
	if userId == "" {
		news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
		if err != nil {
			return []entity.NewsDB{}, err
		}
	} else {
		catdata, err := categoryService.GetCategoryForUser(userId)
		if err != nil {
			news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
			if err != nil {
				return []entity.NewsDB{}, err
			}
		} else {
			l := []string{}
			for _, cat := range catdata {
				l = append(l, cat.CategoryName)
			}
			filter := bson.M{}
			filter["category"] = bson.M{"$in": l}
			news, err = repo.FindSort(filter, bson.M{"publishedAt": -1}, bson.M{}, limit, skip)
			if err != nil || len(news) == 0 {
				news, err = repo.FindSort(bson.M{}, bson.M{"_id": -1}, bson.M{}, limit, skip)
				if err != nil {
					return []entity.NewsDB{}, err
				}
			}

		}
	}
	if err != nil {
		return []entity.NewsDB{}, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(news), func(i, j int) {
		news[i], news[j] = news[j], news[i]
	})
	id, err := primitive.ObjectIDFromHex(newsID)
	if err != nil {
		return news, nil
	}
	var newsByID []entity.NewsDB
	newsByID, err = repo.FindWithIDs(bson.M{"_id": id}, bson.M{})
	if err != nil {
		return news, nil
	}
	news = append(newsByID, news...)
	return news, nil
}
