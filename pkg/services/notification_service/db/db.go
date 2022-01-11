package db

import (
	"MdShorts/pkg/entity"
	notification "MdShorts/pkg/repository/notification_repository"
	"context"
	"errors"
	"log"
	"time"

	category_repository "MdShorts/pkg/repository/category_repository"
	catdb "MdShorts/pkg/services/category_service/dbs"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	opt := option.WithCredentialsFile("md-shorts-firebase-adminsdk-yfgqd-5fe7f1e279.json")
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {

	}
	log.Println("App connect Successfull")

}

var (
	repo = notification.NewNotificationRepository("notification")
)
var (
	categoryService = catdb.NewCategoryService(category_repository.NewCategoryRepository("category"))
)

type notificationService struct{}

func NewNotificationService(repository notification.NotificationRepository) NotificationService {
	repo = repository
	return &notificationService{}
}

func (*notificationService) SendNotificationWithTopic(title, body, topic, userid string) (string, error) {
	var msg entity.MessageData
	msg.ID = primitive.NewObjectID()
	msg.Title = title
	msg.Topic = topic
	msg.Body = body
	msg.UserId = userid
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return "", errors.New("unable to send the notification")
	}
	oneHour := time.Duration(1) * time.Hour

	data := map[string]string{
		"topic":        topic,
		"userid":       userid,
		"Title":        title,
		"Body":         body,
		"click_action": "FLUTTER_NOTIFICATION_CLICK",
	}
	message := &messaging.Message{
		Data: data,
		Android: &messaging.AndroidConfig{
			TTL:      &oneHour,
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Title:       title,
				Body:        body,
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title: title,
						Body:  body,
					},
				},
			},
		},
		Topic: topic,
	}
	_, err = client.Send(ctx, message)
	if err != nil {
		msg.Status = "Failed"
		_, err = repo.InsertOne(msg)
		return "", errors.New("unable to send the notification")
	}
	msg.Status = "Success"
	msg.SentTime = time.Now()
	return repo.InsertOne(msg)
}

func (*notificationService) GetNotifications(limit, skip int, status, userid, topic, title string) ([]entity.MessageData, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if userid != "" {
		les, _ := categoryService.GetCategoryForUser(userid)
		l := []string{}
		for _, cat := range les {
			l = append(l, string(cat.ID.Hex()))
		}
		filter["topic"] = bson.M{"$in": l}
	}
	if topic != "" {
		filter["topic"] = topic
	}
	if title != "" {
		filter["title"] = title
	}

	return repo.FindSort(filter, bson.M{"sentTime": -1}, bson.M{}, 20, 0)
}
