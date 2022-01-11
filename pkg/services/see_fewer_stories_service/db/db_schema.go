package db

import (
	"MdShorts/pkg/entity"
)

type SeeFewerStoriesService interface {
	AddSeeFewerStories(seefewerStories SeeFewerStories) (string, error)
	GetSeeFewerStories(limit, skip int, userid, sourceName string) ([]entity.SeeFewerStoriesDB, error)
}

type SeeFewerStories struct {
	UserId     string `bson:"userId" json:"userId"`
	SourceName string `bson:"sourceName" json:"sourceName"`
}
