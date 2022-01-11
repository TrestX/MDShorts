package db

import (
	"MdShorts/pkg/entity"
	share "MdShorts/pkg/repository/share_repository"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = share.NewShareRepository("share")
)

type shareService struct{}

func NewShareService(repository share.ShareRepository) ShareService {
	repo = repository
	return &shareService{}
}
func (*shareService) AddShare(share Share) (string, error) {
	var shareEntity entity.ShareDB
	shareEntity.SharedTime = time.Now()
	shareEntity.ID = primitive.NewObjectID()
	if share.NewsId == "" {
		return "", errors.New("newsid missing")
	}
	if share.UserId == "" {
		return "", errors.New("userid missing")
	}
	if share.SharedVia == "" {
		return "", errors.New("share via missing")
	}
	shareEntity.NewsId = share.NewsId
	shareEntity.UserId = share.UserId
	shareEntity.SharedVia = share.SharedVia
	return repo.InsertOne(shareEntity)
}

func (*shareService) GetShares(limit, skip int, sharevia, userid, newsid string) ([]entity.ShareDB, error) {
	filter := bson.M{}
	if sharevia != "" {
		filter["shared_via"] = sharevia
	}
	if userid != "" {
		filter["user_id"] = userid
	}
	if newsid != "" {
		filter["news_id"] = newsid
	}
	return repo.Find(filter, bson.M{})
}
