package db

import (
	"MdShorts/pkg/entity"
)

type BookmarkService interface {
	AddBookmark(bookmark BookMark) (string, error)
	UpdateBookmarkStatus(bookmark BookMark, bookmarkid string) (string, error)
	GetBookmarks(limit, skip int, status, userid, newsid string) ([]entity.NewsDB, error)
	GetBookmarksIDs(limit, skip int, status, userid, newsid string) ([]entity.BMNId, error)
}

type BookMark struct {
	UserId string `bson:"user_id" json:"userId,omitempty"`
	NewsId string `bson:"newsId" json:"newsId"`
	Status string `bson:"status" json:"status"`
}
