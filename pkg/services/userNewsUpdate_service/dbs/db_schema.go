package db

import (
	"MdShorts/pkg/entity"
)

type UserNewsUpdateService interface {
	AddNewsUpdateForUser(data entity.UserNewsUpdateDB) (string, error)
	UpdateUserNewsUpdateDB(userId string, currentPage int64, pageHistory []int64) (string, error)
	GetUser(userId string) (entity.UserNewsUpdateDB, error)
}
