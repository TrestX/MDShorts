package db

import (
	"MdShorts/pkg/entity"
	user_news_repository "MdShorts/pkg/repository/user_news_repository"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = user_news_repository.NewUserNewsCheckRepository("usernewscheck")
)

type usersNewsService struct{}

func NewUserNewsService(repository user_news_repository.UserNewsCheckRepository) UserNewsCheckService {
	repo = repository
	return &usersNewsService{}
}
func (*usersNewsService) AddNewsForUser(userId, newsId, status, url, readAt string, timeSpent int64, urlclicked bool) (string, error) {
	var uns entity.UserNewsCheckDB
	uns.ID = primitive.NewObjectID()
	uns.UserID = userId
	uns.NewsID = newsId
	uns.ReadAt = readAt
	uns.AddedAt = time.Now()
	uns.Status = status
	uns.TimeSpentOnReading = timeSpent
	uns.UpdatedAt = time.Now()
	uns.Url = url
	uns.UrlClicked = urlclicked
	_, err := repo.InsertOne(uns)
	if err != nil {
		return "", errors.New("unable to add news status for user")
	}
	return "success", nil
}

func (*usersNewsService) UpdateUserNewsDB(userId, newsId, status, readAt string, timeSpent int64, urlclicked bool) (string, error) {
	if userId == "" {
		return "", errors.New("userid is required")
	}
	if newsId == "" {
		return "", errors.New("newsid is required")
	}

	set := bson.M{}
	set["updated_time"] = time.Now()
	if status != "" {
		set["status"] = status
	}
	set["read_at"] = readAt
	if timeSpent != 0 {
		set["time_spent_on_reading"] = timeSpent
	}
	if urlclicked {
		set["url_clicked"] = true
	}
	set = bson.M{"$set": set}
	filter := bson.M{"user_id": userId, "news_id": newsId}
	return repo.UpdateOne(filter, set)
}

func (*usersNewsService) GetUser(userId, newsId, status string, timeSpent int64, urlclicked bool) ([]entity.UserNewsCheckDB, error) {
	filter := bson.M{}
	if userId != "" {
		filter["user_id"] = userId
	}
	if newsId != "" {
		filter["news_id"] = newsId
	}
	if status != "" {
		filter["status"] = status
	}
	if timeSpent != 0 {
		filter["time_spent"] = timeSpent
	}
	if urlclicked {
		filter["url_clicked"] = urlclicked
	}
	return repo.Find(filter, bson.M{})
}
