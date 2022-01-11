package db

import (
	"MdShorts/pkg/entity"
)

type CategoryService interface {
	AddCategory(categorys Category) (string, error)
	UpdateCategoryStatus(categorys Category, CategoryId string) (string, error)
	GetCategories(limit, skip int, status string) ([]entity.CategoryDB, error)
	GetCategoryWithIDs(categoryId []string) ([]entity.CategoryDB, error)
	GetCategoryForUser(userID string) ([]entity.CategoryDB, error)
}

type Category struct {
	CategoryName string `bson:"category_name" json:"category_name,omitempty"`
	PreSignedUrl string `bson:"pre_signed_url" json:"pre_signed_url"`
	Status       string `bson:"status,omitempty" json:"status,omitempty"`
}
