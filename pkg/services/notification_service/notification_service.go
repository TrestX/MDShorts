package notification_service

import (
	"MdShorts/pkg/repository/notification_repository"
	"MdShorts/pkg/services/notification_service/db"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	notificationService = db.NewNotificationService(notification_repository.NewNotificationRepository("notification"))
)

func SendNotificationWithTopic(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("send notifications", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var message db.Notification
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &message)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send notification"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to send notification"})
		return
	}
	data, err := notificationService.SendNotificationWithTopic(message.Title, message.Body, message.Topic, message.UserId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send notification"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to send notification"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("sendnotification success", logrus.Fields{
		"duration": duration,
	})

}

func Getnotification(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get notifications", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userId := ""
	topic := ""
	title := ""
	status := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	topicS := r.URL.Query().Get("topic")
	statusS := r.URL.Query().Get("status")
	titleS := r.URL.Query().Get("title")
	if userIdS != "" {
		userId = userIdS
	}
	if topicS != "" {
		topic = topicS
	}
	if titleS != "" {
		title = titleS
	}
	if statusS != "" {
		status = statusS
	}
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			limit = 20
		}
	}
	if skipS != "" {
		skip, err = strconv.Atoi(skipS)
		if err != nil {
			skip = 0
		}
	}
	data, err := notificationService.GetNotifications(limit, skip, status, userId, topic, title)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get notification"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get notification"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get notification success", logrus.Fields{
		"duration": duration,
	})
}
