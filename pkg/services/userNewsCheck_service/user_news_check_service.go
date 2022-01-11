package user_news_check_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"MdShorts/pkg/repository/user_news_repository"
	db "MdShorts/pkg/services/userNewsCheck_service/dbs"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	userNewsCheckService = db.NewUserNewsService(user_news_repository.NewUserNewsCheckRepository("usernewscheck"))
)

func AddUserNewsCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("add user news check location", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var userSCheck db.UserCheckSchema
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &userSCheck)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := userNewsCheckService.AddNewsForUser(userSCheck.UserID, userSCheck.NewsID, userSCheck.Status, userSCheck.Url, userSCheck.ReadAt, userSCheck.TimeSpentOnReading, userSCheck.UrlClicked)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add user news check location"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add  user news check location"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("add user news check success", logrus.Fields{
		"duration": duration,
	})
}

func UpdateUserNewsCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update user news check ", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var userSCheck db.UserCheckSchema
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &userSCheck)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := userNewsCheckService.UpdateUserNewsDB(userSCheck.UserID, userSCheck.NewsID, userSCheck.Status, userSCheck.ReadAt, userSCheck.TimeSpentOnReading, userSCheck.UrlClicked)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update user news check "))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update user news check"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("update user news check success", logrus.Fields{
		"duration": duration,
	})
}
