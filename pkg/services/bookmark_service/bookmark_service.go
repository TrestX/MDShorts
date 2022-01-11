package bookmark_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"MdShorts/pkg/entity"
	"MdShorts/pkg/repository/bookmark_repository"
	db "MdShorts/pkg/services/bookmark_service/dbs"

	"github.com/aekam27/trestCommon"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	bookmarkService = db.NewBookmarkService(bookmark_repository.NewBookmarkRepository("bookmark"))
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("adding bookmark", logrus.Fields{"start_time": startTime})
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
	var bookmark db.BookMark
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &bookmark)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := bookmarkService.AddBookmark(bookmark)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to add bookmark"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to add bookmark"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("bookmark added", logrus.Fields{
		"duration": duration,
	})
}

func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("update bookmark status", logrus.Fields{
		"start_time": startTime})
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
	var bookmark db.BookMark
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &bookmark)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	var bookmarkID = mux.Vars(r)["bookmarkId"]
	if bookmarkID == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update bookmark"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update bookmark"})
		return
	}
	data, err := bookmarkService.UpdateBookmarkStatus(bookmark, bookmarkID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to update bookmark"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to update bookmark"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("bookmark updated success", logrus.Fields{
		"duration": duration,
	})
}

func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get bookmark", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userId := ""
	newsId := ""
	status := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	newsIdS := r.URL.Query().Get("newsId")
	statusS := r.URL.Query().Get("status")
	if userIdS != "" {
		userId = userIdS
	}
	if newsIdS != "" {
		newsId = newsIdS
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
	data, err := bookmarkService.GetBookmarks(limit, skip, status, userId, newsId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get bookmarks"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get bookmarks"})
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": []entity.NewsDB{}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get bookmarks success", logrus.Fields{
		"duration": duration,
	})
}

func GetBookmarksNewsIds(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get bookmark newsids", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userId := ""
	newsId := ""
	status := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	newsIdS := r.URL.Query().Get("newsId")
	statusS := r.URL.Query().Get("status")
	if userIdS != "" {
		userId = userIdS
	}
	if newsIdS != "" {
		newsId = newsIdS
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
	data, err := bookmarkService.GetBookmarksIDs(limit, skip, status, userId, newsId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get bookmarks newsids"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get bookmarks newsids"})
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": []entity.NewsDB{}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get bookmarks newsids success", logrus.Fields{
		"duration": duration,
	})
}
