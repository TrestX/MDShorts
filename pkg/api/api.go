package api

import (
	"MdShorts/pkg/entity"
	"encoding/json"
	"strconv"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
)

func GetHealthTopHeadlines(country, language, category string) (entity.TopNewsStruct, error) {
	queryString := "country=" + country + "&pageSize=100&category=" + category + "&language=" + language + "&apiKey=" + viper.GetString("newsapi.key")
	if country == "" {
		queryString = "language=" + language + "&pageSize=100&category=" + category + "&apiKey=" + viper.GetString("newsapi.key")
	}
	url := "https://newsapi.org/v2/top-headlines?" + queryString
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return entity.TopNewsStruct{}, err
	}
	var resp entity.TopNewsStruct
	err = json.Unmarshal(body, &resp)
	return resp, err
}
func GetNewslines(search, time, page string) (entity.TopNewsStruct, error) {
	url := "https://newsapi.org/v2/everything?q=" + search + "&pageSize=100&page=" + page + "&language=en&sortBy=publishedAt&apiKey=" + viper.GetString("newsapi.key")
	body, err := trestCommon.GetApi(" ", url)
	if err != nil {
		return entity.TopNewsStruct{}, err
	}
	var resp entity.TopNewsStruct
	err = json.Unmarshal(body, &resp)
	return resp, err
}

func ClickSend(auth string, number string, otp int) (string, error) {
	strOtp := strconv.Itoa(otp)
	msg := clickSendConstructMsgBody(strOtp, number)
	url := viper.GetString("clicksend.postApi")
	body, err := trestCommon.PostApiwithBasicAuth(auth, url, msg)
	if err != nil {
		return "", err
	}
	var resp entity.ClickSendResponseCode
	err = json.Unmarshal(body, &resp)
	return resp.ResponseCode, err
}

func clickSendConstructMsgBody(otp string, number string) entity.MessageBody {
	var messagestructure []entity.MessageStructure
	var msgbody entity.MessageBody
	var smS entity.MessageStructure
	smS.To = number
	smS.Body = otp + " " + viper.GetString("clicksend.mbody")
	smS.Source = "php"
	messagestructure = append(messagestructure, smS)
	msgbody.Messages = messagestructure
	return msgbody
}
