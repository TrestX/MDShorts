package db

import (
	"MdShorts/pkg/entity"
	user_news_update_repository "MdShorts/pkg/repository/user_news_update_repository"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	repo = user_news_update_repository.NewUserNewsUpdateRepository("usernewsupdate")
)

type usersNewsUpdateService struct{}

func NewUserNewsUpdateService(repository user_news_update_repository.UserNewsUpdateRepository) UserNewsUpdateService {
	repo = repository
	return &usersNewsUpdateService{}
}
func (*usersNewsUpdateService) AddNewsUpdateForUser(data entity.UserNewsUpdateDB) (string, error) {
	return repo.InsertOne(data)
}
func (*usersNewsUpdateService) UpdateUserNewsUpdateDB(userId string, currentPage int64, pageHistory []int64) (string, error) {
	if userId == "" {
		return "", errors.New("userid is required")
	}
	set := bson.M{}
	set["updated_time"] = time.Now()
	if currentPage != 0 {
		set["current_page"] = currentPage
	}
	if len(pageHistory) > 0 {
		set["page_history"] = pageHistory
	}
	set["last_fetched"] = time.Now()
	set = bson.M{"$set": set}
	filter := bson.M{"user_id": userId}
	return repo.UpdateOne(filter, set)
}

func (*usersNewsUpdateService) GetUser(userId string) (entity.UserNewsUpdateDB, error) {
	filter := bson.M{}
	if userId != "" {
		filter["user_id"] = userId
	}
	if userId == "" {
		return entity.UserNewsUpdateDB{}, errors.New("user_id is required")
	}
	return repo.FindOne(filter, bson.M{})
}
