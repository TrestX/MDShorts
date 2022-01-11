package unregistered_user_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	unregistered_user_repository "MdShorts/pkg/repository/unregistered_user_repositroy"
	db "MdShorts/pkg/services/unregistered_user_service/dbs"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	unregisteredUserService = db.NewUnRegisteredUserService(unregistered_user_repository.NewUnRegisteredUserRepository("unregistereduser"))
)

func AddUnregisteredUserService(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("adding un registered user", logrus.Fields{"start_time": startTime})
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
	var unregisteruser db.UnRegisteredUsers
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &unregisteruser)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := unregisteredUserService.AddUnRegisteredUser(unregisteruser)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to  un registered user"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to  un registered user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap(" un registered user added", logrus.Fields{
		"duration": duration,
	})
}

func GetUnregisteredUsers(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("get unregisteredUsers", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	deviceId := ""
	deviceName := ""
	location := ""
	limit := 20
	skip := 0
	var err error
	limitS := r.URL.Query().Get("limit")
	skipS := r.URL.Query().Get("skip")
	deviceIdS := r.URL.Query().Get("deviceId")
	deviceNameS := r.URL.Query().Get("deviceName")
	locationS := r.URL.Query().Get("location")
	if deviceIdS != "" {
		deviceId = deviceIdS
	}
	if deviceNameS != "" {
		deviceName = deviceNameS
	}
	if locationS != "" {
		location = locationS
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
	data, err := unregisteredUserService.GetUnRegisteredUsers(limit, skip, deviceName, location, deviceId)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get unregisteredUsers"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to get unregisteredUsers"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("get unregisteredUsers success", logrus.Fields{
		"duration": duration,
	})
}
