package search_service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"MdShorts/pkg/entity"
	"MdShorts/pkg/repository/search_repository"
	db "MdShorts/pkg/services/search_service/db"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	searchService = db.NewSearchService(search_repository.NewSearchRepository("search"))
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get search", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userId := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	userIdS := r.URL.Query().Get("userId")
	if userIdS != "" {
		userId = userIdS
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
	data, err := searchService.GetSearches(limit, skip, userId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get search"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get search"})
		return
	}
	if data == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": []entity.SearchDB{}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get search success", logrus.Fields{
		"duration": duration,
	})
}
