package util_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

type PreSignedURL struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

func GetPreSignedURL(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("presigned url", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var url PreSignedURL
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &url)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to generate presigned url"})
		return
	}
	data, err := trestCommon.PreSignedUrlAWS(url.Name, url.Path)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to generate presigned url"))
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to generate presigned url"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("presigned url generated", logrus.Fields{
		"duration": duration,
	})
}
