package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopNewsStruct struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"total_results"`
	Articles     []Article `json:"articles"`
}
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}
type NewsDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Category    []string           `bson:"category" json:"category"`
	Url         string             `bson:"url" json:"url"`
	UrlToImage  string             `bson:"urlToImage" json:"urlToImage"`
	Title       string             `bson:"title" json:"title"`
	SourceName  string             `bson:"sourceName" json:"sourceName"`
	PublishedAt time.Time          `bson:"publishedAt" json:"publishedAt"`
	Description string             `bson:"description" json:"description"`
	Author      string             `bson:"author" json:"author"`
	AddedTime   time.Time          `bson:"addedTime" json:"addedTime"`
}
type UserNewsCheckDB struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id"`
	UserID             string             `bson:"user_id" json:"userId"`
	NewsID             string             `bson:"news_id" json:"newsId"`
	TimeSpentOnReading int64              `bson:"time_spent_on_reading" json:"timeSpentOnReading"`
	UrlClicked         bool               `bson:"url_clicked" json:"urlClicked"`
	Status             string             `bson:"status" json:"status"`
	Url                string             `bson:"url" json:"url"`
	ReadAt             string             `bson:"read_at" json:"read_at"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updatedAt"`
	AddedAt            time.Time          `bson:"added_at" json:"addedAt"`
}

type UserNewsUpdateDB struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	PageHistory []int64            `bson:"page_history" json:"page_history"`
	CurrentPage int64              `bson:"current_page" json:"current"`
	LastFetched time.Time          `bson:"last_fetched" json:"last"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}
