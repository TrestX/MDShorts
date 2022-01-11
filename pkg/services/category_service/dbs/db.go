package db

import (
	"MdShorts/pkg/entity"
	"MdShorts/pkg/repository"
	category "MdShorts/pkg/repository/category_repository"
	"MdShorts/pkg/services/profile_service/db"
	"errors"
	"strings"
	"time"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = category.NewCategoryRepository("category")
)
var (
	profileService = db.NewProfileService(repository.NewProfileRepository("users"))
)

type categoryService struct{}

func NewCategoryService(repository category.CategoryRepository) CategoryService {
	repo = repository
	return &categoryService{}
}
func (add *categoryService) AddCategory(category Category) (string, error) {
	var categoryEntity entity.CategoryDB
	categoryEntity.Status = "Active"
	categoryEntity.AddedTime = time.Now()
	categoryEntity.ID = primitive.NewObjectID()
	if category.CategoryName == "" {
		return "", errors.New("category name missing")
	}
	if category.PreSignedUrl == "" {
		return "", errors.New("image url missing")
	}
	categoryEntity.CategoryName = category.CategoryName
	url := createPreSignedDownloadUrl(category.PreSignedUrl)
	categoryEntity.PreSignedUrl = url
	return repo.InsertOne(categoryEntity)
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

func (*categoryService) UpdateCategoryStatus(category Category, categoryid string) (string, error) {
	if categoryid == "" {
		err := errors.New("category id missing")
		trestCommon.ECLog2(
			"update category",
			err,
		)
		return "", err
	}
	id, _ := primitive.ObjectIDFromHex(categoryid)
	_, err := checkByCategoryID(id)
	if err != nil {
		return "", errors.New("invalid category Id")
	}
	setParameters := bson.M{}
	if category.Status != "" {
		setParameters["status"] = category.Status
	}
	if category.CategoryName != "" {
		setParameters["category_name"] = category.CategoryName
	}
	if category.PreSignedUrl != "" {
		setParameters["pre_signed_url"] = createPreSignedDownloadUrl(category.PreSignedUrl)
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	result, err := repo.UpdateOne(filter, set)
	if err != nil {
		trestCommon.ECLog3(
			"update Category ",
			err,
			logrus.Fields{
				"Category_id": categoryid,
			})
		return "", err
	}

	return result, nil
}

func checkByCategoryID(id primitive.ObjectID) (entity.CategoryDB, error) {
	Category, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Category Details section",
			err,
		)
		return Category, err
	}
	return Category, nil
}

func (*categoryService) GetCategories(limit, skip int, status string) ([]entity.CategoryDB, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	category, err := repo.Find(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get category section",
			err,
		)
		return category, err
	}
	// for i := 0; i < len(category); i++ {
	// 	newdownloadurl := createPreSignedDownloadUrl(category[i].PreSignedUrl)
	// 	category[i].PreSignedUrl = newdownloadurl
	// 	setParameters := bson.M{}
	// 	setParameters["pre_signed_url"] = createPreSignedDownloadUrl(newdownloadurl)
	// 	repo.UpdateOne(bson.M{"category_name": category[i].CategoryName}, bson.M{
	// 		"$set": setParameters,
	// 	})
	// }
	return category, nil
}

func (*categoryService) GetCategoryWithIDs(categoryId []string) ([]entity.CategoryDB, error) {
	subFilter := bson.A{}
	for _, item := range categoryId {
		id, _ := primitive.ObjectIDFromHex(item)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter := bson.M{"$or": subFilter}
	category, err := repo.FindWithIDs(filter, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get category section",
			err,
		)
		return category, err
	}
	for i := 0; i < len(category); i++ {
		newdownloadurl := createPreSignedDownloadUrl(category[i].PreSignedUrl)
		category[i].PreSignedUrl = newdownloadurl
	}
	return category, nil
}

func (cat *categoryService) GetCategoryForUser(userID string) ([]entity.CategoryDB, error) {
	var category []entity.CategoryDB
	if userID == "" {
		err := errors.New("UserID missing")
		trestCommon.ECLog2(
			"Get category section",
			err,
		)
		return category, err
	}
	user, _, err := profileService.GetProfile(userID)
	if err != nil {
		trestCommon.ECLog2(
			"Get category section",
			err,
		)
		return category, err
	}

	return cat.GetCategoryWithIDs(user.Category)
}
