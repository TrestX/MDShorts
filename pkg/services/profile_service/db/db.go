package db

import (
	"MdShorts/pkg/entity"
	"MdShorts/pkg/repository"
	"strings"

	"errors"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = repository.NewProfileRepository("users")
)

type profileService struct{}

func NewProfileService(repository repository.ProfileRepository) ProfileService {
	repo = repository
	return &profileService{}
}

func (*profileService) UpdateProfile(profile *Profile, userid string) (string, error) {
	if userid == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	_, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog2(
			"update profile section",
			err,
		)
		return "", err
	}
	if profile.FirstName != "" {
		name := profile.FirstName
		if err != nil {
			trestCommon.ECLog3(
				"update profile section",
				err,
				logrus.Fields{
					"user_id":      userid,
					"profile.name": profile.FirstName,
				})
			return "", err
		}
		setParameters["first_name"] = name
	}
	if profile.LastName != "" {
		name := profile.LastName
		if err != nil {
			trestCommon.ECLog3(
				"update profile section",
				err,
				logrus.Fields{
					"user_id":      userid,
					"profile.name": profile.LastName,
				})
			return "", err
		}
		setParameters["last_name"] = name
	}
	if profile.PhoneNo != "" {
		phoneNo, err := trestCommon.Encrypt(profile.PhoneNo)
		if err != nil {
			trestCommon.ECLog3(
				"update profile section",
				err,
				logrus.Fields{
					"user_id":         userid,
					"profile.phoneNo": profile.PhoneNo,
				})
			return "", err
		}
		setParameters["phone_no"] = phoneNo
	}
	if profile.Designation != "" {

		setParameters["designation"] = profile.Designation
	}
	if len(profile.Speciality) > 0 {
		setParameters["speciality"] = profile.Speciality
	}
	if len(profile.Categories) > 0 {
		setParameters["categories"] = profile.Categories
	}
	if profile.Status != "" {
		setParameters["status"] = profile.Status
	}
	if profile.Address != "" {
		setParameters["address.address"] = profile.Address
		if err != nil {
			trestCommon.ECLog3(
				"update profile section",
				err,
				logrus.Fields{
					"user_id": userid,
					"address": profile.Address,
				})
			return "", err
		}
	}
	if profile.State != "" {
		setParameters["address.state"] = profile.State
	}
	if profile.City != "" {
		setParameters["address.city"] = profile.City
	}
	if profile.Country != "" {
		setParameters["address.country"] = profile.Country
	}
	if profile.Pin != "" {
		setParameters["address.pin"] = profile.Pin
	}
	if profile.UrlToProfileImage != "" {
		setParameters["url_to_profile_image"] = profile.UrlToProfileImage
	}
	if profile.About != "" {
		setParameters["about"] = profile.About
	}
	setParameters["update_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile section",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}

	return result, nil
}

func (*profileService) GetProfile(userID string) (entity.ProfileDB, int64, error) {
	if userID == "" {
		err := errors.New("user id missing")
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return entity.ProfileDB{}, 0, err
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"GetProfile section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.ProfileDB{}, 0, err
	}
	profile, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"GetProfile section",
			err,
		)
		return profile, 0, err
	}
	// profile.Password = ""
	newUrl := createPreSignedDownloadUrl(profile.UrlToProfileImage)
	profile.UrlToProfileImage = newUrl
	per := calculatePercentage(profile)
	if len(profile.Speciality) < 1 {
		profile.Speciality = []string{""}
	}
	return profile, per, nil
}

func calculatePercentage(profile entity.ProfileDB) int64 {
	initial := 12
	if profile.FirstName != "" && profile.LastName != "" {
		initial = initial + 11
	}
	if profile.About != "" {
		initial = initial + 11
	}
	if profile.Designation != "" {
		initial = initial + 11
	}
	if profile.PhoneNo != "" {
		initial = initial + 11
	}
	if len(profile.Speciality) > 0 {
		initial = initial + 11
	}
	if profile.UrlToProfileImage != "" {
		initial = initial + 11
	}
	if profile.UrlToProfileImage != "" {
		initial = initial + 11
	}
	if profile.Address.Address != "" {
		initial = initial + 3
	}
	if profile.Address.City != "" {
		initial = initial + 2
	}
	if profile.Address.Country != "" {
		initial = initial + 2
	}
	if profile.Address.Pin != "" {
		initial = initial + 2
	}
	if profile.Address.State != "" {
		initial = initial + 2
	}
	return int64(initial)
}
func createPreSignedDownloadUrl(url string) string {
	s := strings.Split(url, "?")
	if len(s) > 0 {
		o := strings.Split(s[0], "/")
		if len(o) > 3 {
			fileName := o[4]
			path := o[3]
			downUrl, _ := trestCommon.PreSignedDownloadUrlAWS(fileName, path)
			return downUrl
		}
	}
	return ""
}
func checkUser(userID string) (entity.ProfileDB, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		trestCommon.ECLog3(
			"CheckUser section",
			err,
			logrus.Fields{
				"user_id": userID,
			},
		)
		return entity.ProfileDB{}, err
	}
	return repo.FindOne(bson.M{"_id": id}, bson.M{})
}

func (*profileService) ChangePassword(profile *Profile, userid string) (string, error) {
	// if profile.Password == "" {
	// 	trestCommon.ECLog3("new password", errors.New("there was an error while updating the password"), logrus.Fields{"email": profile.Email})
	// 	return "", errors.New("there was an error while updating the password")
	// }
	salt := viper.GetString("salt")
	hash, err := bcrypt.GenerateFromPassword([]byte(profile.Password+salt), 5)
	if err != nil {
		trestCommon.ECLog3("hash new password", err, logrus.Fields{"email": profile.Email})
		return "", err
	}
	password := string(hash)
	id, _ := primitive.ObjectIDFromHex(userid)
	setParameters := bson.M{}
	userData, err := checkUser(userid)
	if err != nil {
		trestCommon.ECLog3("hash new password", err, logrus.Fields{"email": profile.Email})
		return "", err
	}
	if userData.Email != profile.Email {
		trestCommon.ECLog3("new password", errors.New("there was an error while updating the password"), logrus.Fields{"email": profile.Email})
		return "", errors.New("there was an error while updating the password")
	}
	if password != "" {
		setParameters["password"] = password
	}
	setParameters["update_time"] = time.Now()
	setParameters["password_updated_on_time"] = time.Now()
	setParameters["last_login_device_info"] = profile.LastLoginDeviceInfo
	setParameters["last_login_location"] = profile.LastLoginLocation
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	_, err = repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update profile password",
			err,
			logrus.Fields{
				"user_id": userid,
				"profile": profile,
			})
		return "", err
	}
	_, err = sendConfirmationEmail(userData.Email)
	if err != nil {
		trestCommon.ECLog3("hashAndInsertData Insert failed", err, logrus.Fields{"email": profile.Email})
	}
	return "password changed successfully", nil
}

func sendConfirmationEmail(email string) (string, error) {
	emailSentTime := time.Now()
	verificationCode := trestCommon.GetRandomString(16)
	sendCode, err := trestCommon.Encrypt(email + ":" + verificationCode)
	if err != nil {
		trestCommon.ECLog2("send verification email encryption failed", err)
		return "", err
	}
	_, err = trestCommon.SendPasswordConfirmation(email, sendCode)
	if err != nil {
		trestCommon.ECLog2("send verification email failed", err)
		return "", err
	}
	_, err = repo.UpdateOne(bson.M{"email": email}, bson.M{"$set": bson.M{"password_reset_confirmation_email_sent_time": emailSentTime}})
	if err != nil {
		trestCommon.ECLog2("send verification email update failed", err)
		return "", err
	}
	return "email sent successfully", nil
}
