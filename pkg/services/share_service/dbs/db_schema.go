package db

import (
	"MdShorts/pkg/entity"
)

type ShareService interface {
	AddShare(share Share) (string, error)
	GetShares(limit, skip int, sharevia, userid, newsid string) ([]entity.ShareDB, error)
}

type Share struct {
	UserId    string `bson:"user_id" json:"userId,omitempty"`
	NewsId    string `bson:"newsId" json:"newsId"`
	SharedVia string `bson:"shared_via" json:"sharedVia"`
}
