package db

import (
	"MdShorts/pkg/entity"
)

type UserNewsCheckService interface {
	AddNewsForUser(userId, newsId, status, url, readAt string, timeSpent int64, urlclicked bool) (string, error)
	UpdateUserNewsDB(userId, newsId, status, readAt string, timeSpent int64, urlclicked bool) (string, error)
	GetUser(userId, newsId, status string, timeSpent int64, urlclicked bool) ([]entity.UserNewsCheckDB, error)
}

type UserCheckSchema struct {
	UserID             string `bson:"user_id" json:"userId,omitempty"`
	NewsID             string `bson:"news_id" json:"newsId,omitempty"`
	TimeSpentOnReading int64  `bson:"time_spent_on_reading" json:"timeSpentOnReading"`
	UrlClicked         bool   `bson:"url_clicked" json:"urlClicked,omitempty"`
	Status             string `bson:"status" json:"status,omitempty"`
	Url                string `bson:"url" json:"url,omitempty"`
	ReadAt             string `bson:"read_at" json:"read_at,omitempty"`
}
