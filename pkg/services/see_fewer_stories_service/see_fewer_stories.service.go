package see_fewer_stories_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	see_fewer_stories "MdShorts/pkg/repository/see_fewer_stories"
	db "MdShorts/pkg/services/see_fewer_stories_service/db"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	seeFewerStoriesService = db.NewseeFewerStoriesService(see_fewer_stories.NewSeeFewerStoriesRepository("see_fewer_stories"))
)

func AddSeeFewerStories(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("adding see fewer stories", logrus.Fields{"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	_, err := trestCommon.DecodeToken(tokenString[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var seeFewerStories db.SeeFewerStories
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &seeFewerStories)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := seeFewerStoriesService.AddSeeFewerStories(seeFewerStories)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add seeFewerStories"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add seeFewerStories"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("seeFewerStories added", logrus.Fields{
		"duration": duration,
	})
}

func GetShares(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get share", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userId := ""
	sourceName := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	sourceNameS := r.URL.Query().Get("sourceName")
	if userIdS != "" {
		userId = userIdS
	}
	if sourceNameS != "" {
		sourceName = sourceNameS
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
	data, err := seeFewerStoriesService.GetSeeFewerStories(limit, skip, userId, sourceName)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get seeFewerStories"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get seeFewerStories"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get seeFewerStories success", logrus.Fields{
		"duration": duration,
	})
}
