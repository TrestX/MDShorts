package db

import "MdShorts/pkg/entity"

type ProfileService interface {
	UpdateProfile(profile *Profile, userid string) (string, error)
	GetProfile(userID string) (entity.ProfileDB, int64, error)
	ChangePassword(profile *Profile, userid string) (string, error)
}

type Profile struct {
	Status              string      `bson:"status,omitempty" json:"status,omitempty"`
	FirstName           string      `bson:"first_name" json:"firstName"`
	LastLoginDeviceInfo interface{} `bson:"last_login_device_info" json:"lastLoginDeviceInfo"`
	LastLoginLocation   string      `bson:"last_login_location" json:"lastLoginLocation"`
	Email               string      `bson:"email" json:"email"`
	Password            string      `bson:"password" json:"password"`
	LastName            string      `bson:"last_name" json:"lastName"`
	PhoneNo             string      `bson:"phone_no" json:"phoneNumber"`
	CountryCode         string      `bson:"country_code" json:"countryCode"`
	Designation         string      `bson:"designation" json:"designation"`
	Speciality          []string    `bson:"speciality" json:"speciality"`
	Categories          []string    `bson:"categories" json:"categories"`
	Address             string      `bson:"address" json:"address,omitempty"`
	Country             string      `bson:"country" json:"country,omitempty"`
	Pin                 string      `bson:"pin" json:"pin,omitempty"`
	City                string      `bson:"city" json:"city,omitempty"`
	State               string      `bson:"state" json:"state,omitempty"`
	About               string      `bson:"about" json:"about,omitempty"`
	UrlToProfileImage   string      `bson:"url_to_profile_image" json:"urlToProfileImage,omitempty"`
}
