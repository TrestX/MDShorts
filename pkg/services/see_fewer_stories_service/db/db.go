package db

import (
	"MdShorts/pkg/entity"
	see_fewer_stories "MdShorts/pkg/repository/see_fewer_stories"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = see_fewer_stories.NewSeeFewerStoriesRepository("see_fewer_stories")
)

type seeFewerStoriesService struct{}

func NewseeFewerStoriesService(repository see_fewer_stories.SeeFewerStoriesRepository) SeeFewerStoriesService {
	repo = repository
	return &seeFewerStoriesService{}
}
func (*seeFewerStoriesService) AddSeeFewerStories(seefewerStories SeeFewerStories) (string, error) {
	var seeFewerStorie entity.SeeFewerStoriesDB
	seeFewerStorie.AddedTime = time.Now()
	seeFewerStorie.ID = primitive.NewObjectID()
	if seefewerStories.UserId == "" {
		return "", errors.New("userid missing")
	}
	if seefewerStories.SourceName == "" {
		return "", errors.New("source name missing")
	}
	seeFewerStorie.SourceName = seefewerStories.SourceName
	seeFewerStorie.UserId = seefewerStories.UserId
	return repo.InsertOne(seeFewerStorie)
}

func (*seeFewerStoriesService) GetSeeFewerStories(limit, skip int, userid, sourceName string) ([]entity.SeeFewerStoriesDB, error) {
	filter := bson.M{}
	if userid != "" {
		filter["userId"] = userid
	}
	if sourceName != "" {
		filter["sourceName"] = sourceName
	}
	return repo.Find(filter, bson.M{})
}
