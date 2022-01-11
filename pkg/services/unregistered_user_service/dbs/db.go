package db

import (
	"MdShorts/pkg/entity"
	unregistered_user "MdShorts/pkg/repository/unregistered_user_repositroy"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = unregistered_user.NewUnRegisteredUserRepository("unregistereduser")
)

type unRegisteredUserService struct{}

func NewUnRegisteredUserService(repository unregistered_user.UnRegisteredUserRepository) UnRegisteredUserService {
	repo = repository
	return &unRegisteredUserService{}
}
func (*unRegisteredUserService) AddUnRegisteredUser(unregistereduser UnRegisteredUsers) (string, error) {
	var unregisteredUserEntity entity.UnRegisteredUsersDB
	unregisteredUserEntity.Time = time.Now()
	unregisteredUserEntity.ID = primitive.NewObjectID()
	if unregistereduser.DeviceID == "" {
		return "", errors.New("device id missing")
	}
	unregisteredUserEntity.DeviceID = unregistereduser.DeviceID
	unregisteredUserEntity.DeviceName = unregistereduser.DeviceName
	unregisteredUserEntity.Location = unregistereduser.Location
	return repo.InsertOne(unregisteredUserEntity)
}

func (*unRegisteredUserService) GetUnRegisteredUsers(limit, skip int, devicename, location, deviceid string) ([]entity.UnRegisteredUsersDB, error) {
	filter := bson.M{}
	if devicename != "" {
		filter["device_name"] = devicename
	}
	if deviceid != "" {
		filter["device_id"] = deviceid
	}
	if location != "" {
		filter["location"] = location
	}
	return repo.Find(filter, bson.M{})
}
