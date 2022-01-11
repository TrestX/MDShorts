package db

import (
	"MdShorts/pkg/entity"
)

type UnRegisteredUserService interface {
	AddUnRegisteredUser(share UnRegisteredUsers) (string, error)
	GetUnRegisteredUsers(limit, skip int, devicename, location, deviceid string) ([]entity.UnRegisteredUsersDB, error)
}

type UnRegisteredUsers struct {
	DeviceID   string `bson:"device_id" json:"DeviceID"`
	DeviceName string `bson:"device_name" json:"DeviceName"`
	Location   string `bson:"location" json:"Location"`
}
