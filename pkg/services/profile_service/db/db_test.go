package db_test

import (
	"MdShorts/pkg/entity"
	db "MdShorts/pkg/services/profile_service/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	useriddoesnotexist = "608f99f5d9f3f6d0728eebc3"
	email              = "test@test.com"
	name               = "test"
	lname              = "tset"
	phoneno            = "9087678123981"
	Designation        = "main"
	account_type       = "ERP"
)

var userid = primitive.NewObjectID()

var profileDB = entity.ProfileDB{
	ID:    userid,
	Email: "test@test.org",
	// Password:         "$2a$05$FbQfv6Z0/9tbF0w/jWwfDOLjAK157EB0SU0MjjnDctVHtOmDE7/g.",
	Status:           "verified",
	VerificationCode: "9h2BdTI0WwmGH1Gl",
	// TermsChecked:     true,
	Designation: "Test",
	FirstName:   "Te",
	LastName:    "sT",
	PhoneNo:     "9876543210",
	Address:     address,
}
var address = entity.AddressDB{
	Address: "aaaaa",
	City:    "city",
	State:   "state",
	Pin:     "pin",
	Country: "country",
}

type TESTPROFILE struct {
	Status bool   `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	Token  string `json:"token,omitempty"`
}

var profile = db.Profile{
	FirstName:   name,
	LastName:    lname,
	PhoneNo:     phoneno,
	Designation: Designation,
	Address:     "aaaaa",
	City:        "city",
	State:       "state",
	Pin:         "pin",
	Country:     "country",
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) UpdateOne(filter, update bson.M) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}

func (mock *MockRepository) FindOne(filter, projection bson.M) (entity.ProfileDB, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.ProfileDB), args.Error(1)
}
func (mock *MockRepository) DeleteOne(filter bson.M) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) Find(filter, projection bson.M) ([]entity.ProfileDB, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.ProfileDB), args.Error(1)
}
func (mock *MockRepository) InsertOne(document interface{}) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}

func TestGetProfile(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("FindOne").Return(profileDB, nil)
	testService := db.NewProfileService(mockRepo)

	t.Run("User Exist", func(t *testing.T) {
		newProfile, _, _ := testService.GetProfile(userid.Hex())
		assert.Equal(t, newProfile, profileDB)
	})
	t.Run("User ID Missing", func(t *testing.T) {
		_, _, err := testService.GetProfile("")
		assert.Equal(t, "user id missing", err.Error())
	})
	t.Run("User not in DB", func(t *testing.T) {
		_, _, err := testService.GetProfile("sjasjkaskhjk")
		assert.NotNil(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)

	mockRepo.On("UpdateOne").Return("Updated Successfully", nil)
	testService := db.NewProfileService(mockRepo)
	t.Run("Update Successful", func(t *testing.T) {
		var profile *db.Profile = &db.Profile{}
		profile.FirstName = name
		profile.PhoneNo = phoneno
		got, _ := testService.UpdateProfile(profile, userid.Hex())
		want := "Updated Successfully"
		assert.Equal(t, got, want)
	})
	t.Run("Update For no user", func(t *testing.T) {
		var profile *db.Profile = &db.Profile{}
		profile.FirstName = name
		profile.PhoneNo = phoneno
		_, err := testService.UpdateProfile(profile, "")
		assert.Equal(t, "user id missing", err.Error())
	})
}
