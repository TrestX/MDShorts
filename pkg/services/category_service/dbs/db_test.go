package db_test

import (
	"MdShorts/pkg/entity"
	db "MdShorts/pkg/services/category_service/dbs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var catid = primitive.NewObjectID()
var categoryDB = []entity.CategoryDB{{
	ID:           catid,
	CategoryName: "abcd",
	Status:       "Active",
	PreSignedUrl: ""},
}
var category = db.Category{
	CategoryName: "abcd",
	Status:       "Active",
	PreSignedUrl: "http://localhost",
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) UpdateOne(filter, update bson.M) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}

func (mock *MockRepository) FindOne(filter, projection bson.M) (entity.CategoryDB, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.CategoryDB), args.Error(1)
}
func (mock *MockRepository) DeleteOne(filter bson.M) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) FindWithIDs(filter, projection bson.M) ([]entity.CategoryDB, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.CategoryDB), args.Error(1)
}
func (mock *MockRepository) Find(filter, projection bson.M) ([]entity.CategoryDB, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.CategoryDB), args.Error(1)
}
func (mock *MockRepository) InsertOne(document entity.CategoryDB) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
func TestSetCategory(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Find").Return(categoryDB, nil)
	mockRepo.On("InsertOne").Return("Added Successfully", nil)
	testService := db.NewCategoryService(mockRepo)
	t.Run("Set Category Successful", func(t *testing.T) {
		got, _ := testService.AddCategory(category)
		want := "Added Successfully"
		assert.Equal(t, got, want)
	})
	t.Run("Set Category Name not Added", func(t *testing.T) {
		category.CategoryName = ""
		_, err := testService.AddCategory(category)

		assert.Equal(t, "category name missing", err.Error())
	})
	t.Run("Set Category With url", func(t *testing.T) {
		_, err := testService.AddCategory(category)
		want := "image url is missing"
		assert.Equal(t, want, err.Error())
	})
}
func TestGetProfile(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("FindOne").Return(categoryDB, nil)
	testService := db.NewCategoryService(mockRepo)

	t.Run("Category Exist", func(t *testing.T) {
		got, _ := testService.GetCategories(10, 10, "active")
		assert.Equal(t, got, categoryDB)
	})
}
