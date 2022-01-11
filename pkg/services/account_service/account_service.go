package account_service

import (
	"MdShorts/pkg/repository"
	db "MdShorts/pkg/services/account_service/dbs"
	"strconv"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	accountService = db.NewSignUpService(repository.NewProfileRepository("users"))
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("sign up email sent", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, err := GetCredentials(r)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to get credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return

	}
	data, err := accountService.SignUp(user)
	if err != nil || data == "" {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to sent singup email"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "email already registered or phone number"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": "sign up email sent successfully", "token": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("sign up email sent successfully", logrus.Fields{"duration": duration})
}

func Login(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("login email sent", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, err := GetCredentials(r)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to parse credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return

	}
	data, err := accountService.Login(user)
	if err != nil {
		if err.Error() == "user not verified" {
			trestCommon.ECLog1(errors.Wrapf(err, "unable to login"))
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Email not Verified"})
			return
		}
		trestCommon.ECLog1(errors.Wrapf(err, "unable to login"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "invalid credentials"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "token": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login email sent successfully", logrus.Fields{"duration": duration})
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Verify Email", logrus.Fields{
		"start_time": startTime})
	var cred db.Credentials
	verificationCode := mux.Vars(r)["code"]
	plainCode, _ := trestCommon.Decrypt(verificationCode)

	cred.Email = strings.Split(plainCode, ":")[0]
	cred.VerificationCode = strings.Split(plainCode, ":")[1]
	message, err := accountService.VerifyEmail(cred)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify email"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": message})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("Email verified", logrus.Fields{"duration": duration})
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Verify OTP", logrus.Fields{
		"start_time": startTime})
	var otp db.OTP
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &otp)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify mobile number"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	message, err := accountService.VerifyOTP(otp)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify mobile number"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": message})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("MObile number verified", logrus.Fields{"duration": duration})
}

func EmailVerifyOTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Email Verify OTP Login", logrus.Fields{
		"start_time": startTime})
	var creds db.Credentials
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &creds)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify email"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	token, data, per, err := accountService.VerifyEmailOTP(creds)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify email"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "token": token, "data": data, "profilePer": strconv.Itoa(int(per))})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login Success", logrus.Fields{"duration": duration})
}

func ResendOTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Resend OTP", logrus.Fields{
		"start_time": startTime})
	var otp db.OTP
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &otp)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to resend otp on mobile number"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	message, err := accountService.ResendOTP(otp)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to resend otp on mobile number"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": message})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("OTP Send", logrus.Fields{"duration": duration})
}

func SocialMedialogin(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("social media login", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, err := GetCredentials(r)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to parse credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	token, data, per, err := accountService.SocialMedialogin(user)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify email"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "token": token, "data": data, "profilePer": strconv.Itoa(int(per))})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login successfully", logrus.Fields{"duration": duration})
}

func SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Send Email", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	claims, err := trestCommon.DecodeToken(strings.Split(r.Header.Get("Authorization"), " ")[1])
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate token"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "authorization failed"})
		return
	}
	var emailVerification db.EmailVerification
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &emailVerification)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "failed to authenticate payload"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "invalid payload"})
		return
	}
	data, err := accountService.SendVerificationEmail(emailVerification.Email, claims["email"].(string), claims["userid"].(string))
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send verification email"))

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	if data == "email sent successfully" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": data, "token": ""})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": "success", "token": data})
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("Email sent", logrus.Fields{"duration": duration})
}
func SendPasswordResetOtp(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Send Password reset link", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user, err := GetCredentials(r)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to parse credentials"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := accountService.SendResetOTP(user.Email)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to send password reset otp"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("Reset link sent", logrus.Fields{"duration": duration})
}

func VerifyPasswordResetLink(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("Verify password reset link", logrus.Fields{
		"start_time": startTime})
	var cred db.Credentials
	verificationCode := mux.Vars(r)["code"]
	plainCode, _ := trestCommon.Decrypt(verificationCode)

	cred.Email = strings.Split(plainCode, ":")[0]
	cred.PasswordResetCode = strings.Split(plainCode, ":")[1]
	message, email, err := accountService.VerifyResetLink(cred)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to verify the Password reset link"))
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": message, "email": email})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("Reset link verified", logrus.Fields{"duration": duration})
}

// func UpdatePassword(w http.ResponseWriter, r *http.Request) {
// 	startTime := time.Now()
// 	trestCommon.DLogMap("update password", logrus.Fields{
// 		"start_time": startTime})
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	user, err := GetCredentials(r)
// 	if err != nil {
// 		trestCommon.ECLog1(errors.Wrapf(err, "unable to get credentials"))
// 		w.WriteHeader(http.StatusUnsupportedMediaType)
// 		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
// 		return
// 	}
// 	data, err := accountService.UpdatePassword(user)
// 	if err != nil || data == "" {
// 		trestCommon.ECLog1(errors.Wrapf(err, "unable to Update the Password"))
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "something went wrong"})
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "message": data})
// 	endTime := time.Now()
// 	duration := endTime.Sub(startTime)
// 	trestCommon.DLogMap("Password Updated", logrus.Fields{"duration": duration})
// }

func GetCredentials(r *http.Request) (db.Credentials, error) {
	var user db.Credentials

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}
	user.Email = strings.TrimSpace(user.Email)
	return user, err
}
