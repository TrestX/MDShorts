package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileDB struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	Email             string             `bson:"email" json:"email"`
	Status            string             `bson:"status" json:"status"`
	FirstName         string             `bson:"first_name" json:"firstName"`
	LastName          string             `bson:"last_name" json:"lastName"`
	PhoneNo           string             `bson:"phone_no" json:"phoneNumber"`
	Designation       string             `bson:"designation" json:"designation"`
	Speciality        []string           `bson:"speciality" json:"speciality"`
	Address           AddressDB          `bson:"address" json:"address"`
	About             string             `bson:"about" json:"about"`
	UrlToProfileImage string             `bson:"url_to_profile_image" json:"urlToProfileImage"`
	Category          []string           `bson:"categories" json:"categories"`
	// TermsChecked        bool               `bson:"terms_and_condition" json:"termsAndConditions"`
	// Password            string             `bson:"password" json:"password"`
	CreatedTime         time.Time `bson:"created_time" json:"createdTime"`
	EmailLoginOTP       string    `bson:"email_login_otp" json:"emailLoginOtp"`
	OTP                 string    `bson:"otp_code" json:"otp_code"`
	UpdateTime          time.Time `bson:"update_time" json:"updateTime"`
	EmailSentTime       time.Time `bson:"email_sent_time" json:"emailSentTime"`
	VerificationCode    string    `bson:"verification_code" json:"verificationCode"`
	PasswordResetCode   string    `bson:"password_reset_code" json:"passwordResetCode"`
	CountryCode         string    `bson:"country_code" json:"countryCode"`
	PasswordResetTime   time.Time `bson:"password_reset_time" json:"passwordResetTime"`
	LastLoginDeviceID   string    `bson:"last_login_device_id" json:"lastLoginDeviceID"`
	LastLoginDeviceName string    `bson:"last_login_device_name" json:"lastLoginDeviceName"`
	LastLoginLocation   string    `bson:"last_login_location" json:"lastLoginLocation"`
}

type AddressDB struct {
	Address string `bson:"address" json:"address"`
	Country string `bson:"country" json:"country"`
	Pin     string `bson:"pin" json:"pin"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
}

type UnRegisteredUsersDB struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	DeviceID   string             `bson:"device_id" json:"DeviceID"`
	DeviceName string             `bson:"device_name" json:"DeviceName"`
	Location   string             `bson:"location" json:"Location"`
	Time       time.Time          `bson:"time" json:"Time"`
}
