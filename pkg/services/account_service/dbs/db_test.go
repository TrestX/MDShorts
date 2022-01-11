package db_test

import (
	"MdShorts/pkg/entity"
	db "MdShorts/pkg/services/account_service/dbs"

	"errors"
	"strings"
	"testing"

	"github.com/aekam27/trestCommon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/bson"
)

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

const email = "support@wiredsoft.org"

var token string

var profileDB = entity.ProfileDB{
	Email: "test@test.org",
	// Password:          "$2a$05$FwMkphzDSa10Qyrsg0xPX.VFfLsN19IKym9Ro/K7bgzareHvfPk9e",
	Status:           "created",
	VerificationCode: "9h2BdTI0WwmGH1Gl",
	// TermsChecked:      true,
	Designation:       "Test",
	FirstName:         "Te",
	LastName:          "st",
	PhoneNo:           "9876543210",
	PasswordResetCode: "9h2BdTI0WwmGH1Gl",
}
var profileDBVerified = entity.ProfileDB{
	Email: "test@test.org",
	// Password:          "$2a$05$FwMkphzDSa10Qyrsg0xPX.VFfLsN19IKym9Ro/K7bgzareHvfPk9e",
	Status:           "verified",
	VerificationCode: "9h2BdTI0WwmGH1Gl",
	// TermsChecked:      true,
	Designation:       "Test",
	FirstName:         "Te",
	LastName:          "st",
	PhoneNo:           "9876543210",
	PasswordResetCode: "9h2BdTI0WwmGH1Gl",
}

func TestSignUp(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("InsertOne").Return("Added Successfully", nil)
	mockRepo.On("FindOne").Return(profileDB, errors.New("mongo: no documents in result"))
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)
	t.Run("SignUp Successfull", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = email

		got, err := testService.SignUp(cred)
		assert.Nil(t, err)
		assert.NotEmpty(t, got)
		token = got
	})

	t.Run("SignUp Wrong Email format", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "abcac.com"

		_, err := testService.SignUp(cred)
		assert.NotNil(t, err)
		want := "invalid email"
		assert.Equal(t, want, err.Error())
	})
	t.Run("SignUp Missing Email", func(t *testing.T) {
		cred := db.Credentials{}
		_, err := testService.SignUp(cred)
		assert.NotNil(t, err)
		want := "email missing"
		assert.Equal(t, want, err.Error())
	})
	t.Run("SignUp Missing password", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "test@test.com"
		_, err := testService.SignUp(cred)
		assert.NotNil(t, err)
		want := "password missing"
		assert.Equal(t, want, err.Error())
	})
	t.Run("SignUp Failed user already exists", func(t *testing.T) {
		cred := db.Credentials{}
		mockRepo := new(MockRepository)
		mockRepo.On("InsertOne").Return("Added Successfully", nil)
		mockRepo.On("FindOne").Return(profileDB, nil)
		testService := db.NewSignUpService(mockRepo)
		cred.Email = "test@test.org"
		got, err := testService.SignUp(cred)
		assert.Empty(t, got)
		assert.Equal(t, "email already registed", err.Error())
	})
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(profileDBVerified, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)

	t.Run("Login Success", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "test@test.org"
		got, err := testService.Login(cred)
		assert.Nil(t, err)
		assert.NotEmpty(t, got)

	})

	mockRepo = new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService = db.NewSignUpService(mockRepo)
	t.Run("Login Failed User Not Verified", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "test@test.org"
		got, err := testService.Login(cred)
		assert.NotNil(t, err)
		assert.Empty(t, got)

	})

	t.Run("Login Failed Wrong password", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "test@test.org"
		got, err := testService.Login(cred)
		assert.NotNil(t, err)
		assert.Empty(t, got)
	})
	t.Run("Login Failed Wrong email", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "em@email.com"
		got, err := testService.Login(cred)
		assert.NotNil(t, err)
		assert.Empty(t, got)

		assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())
	})
	t.Run("Login Failed Empty Email", func(t *testing.T) {
		cred := db.Credentials{}
		got, err := testService.Login(cred)
		assert.NotNil(t, err)
		assert.Empty(t, got)
		want := "email missing"
		assert.Equal(t, want, err.Error())
	})
	t.Run("Login Failed Empty Password", func(t *testing.T) {
		cred := db.Credentials{}
		cred.Email = "em@email.com"
		got, err := testService.Login(cred)
		assert.NotNil(t, err)
		assert.Empty(t, got)
		assert.Equal(t, "password missing", err.Error())
	})
}

func TestSendVerificationEmail(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)
	t.Run("Send Email with token", func(t *testing.T) {

		got, err := testService.SendVerificationEmail(email, "", "")
		assert.Nil(t, err)
		want := "email sent successfully"
		assert.Equal(t, got, want)
	})
	t.Run("Send Email without email", func(t *testing.T) {
		got, err := testService.SendVerificationEmail("", "", "")
		assert.NotNil(t, err)
		want := "email sent successfully"
		assert.NotEqual(t, got, want)
	})
}
func TestSendResetLink(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)
	t.Run("Send Email with token", func(t *testing.T) {

		got, err := testService.SendResetLink("support@wiredsoft.org")
		assert.Nil(t, err)
		want := "Reset link sent successfully"
		assert.Equal(t, got, want)
	})
	t.Run("Send Email without email", func(t *testing.T) {
		_, err := testService.SendResetLink("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "email missing")
	})
	t.Run("Send Email without valid email", func(t *testing.T) {
		_, err := testService.SendResetLink("fsefewfe")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid email")
	})
	mockRepo = new(MockRepository)
	mockRepo.On("FindOne").Return(entity.ProfileDB{}, errors.New("user doesnot exist"))
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService = db.NewSignUpService(mockRepo)
	t.Run("User doesnot exist", func(t *testing.T) {
		_, err := testService.SendResetLink(email)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user doesnot exist")
	})
	mockRepo = new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("", errors.New("error updating password reset code"))
	testService = db.NewSignUpService(mockRepo)
	t.Run("Error While updating Password reset code in db", func(t *testing.T) {
		_, err := testService.SendResetLink(email)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error updating password reset code")
	})
}
func TestVerifyUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(entity.ProfileDB{}, errors.New("user doesn't exist"))
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)
	t.Run("Verfiy User Wrong Email", func(t *testing.T) {

		cred := db.Credentials{}
		cred.Email = "neeraj1@wiredsoft.org"
		cred.VerificationCode = profileDB.VerificationCode
		_, err := testService.VerifyEmail(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "user doesn't exist", err.Error())

	})
	mockRepo = new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService = db.NewSignUpService(mockRepo)
	t.Run("Verfiy User Wrong code", func(t *testing.T) {

		cred := db.Credentials{}
		cred.Email = profileDB.Email
		cred.VerificationCode = "sdmklmdkljslakj"
		_, err := testService.VerifyEmail(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "unauthorized user", err.Error())

	})
	t.Run("Verfiy User Success", func(t *testing.T) {
		cred := db.Credentials{}
		credentials, err := trestCommon.Decrypt("86c1e48f56bd8231f1be1f3946b7e7357991ba323f813957b68af0e2c1fd4df4a94349eea31f30e781343c9ddb6446b1c0b3c0879e5ad4acbc2c6e6a00cdb2705a")
		assert.Nil(t, err)
		cred.VerificationCode = strings.Split(credentials, ":")[1]

		cred.Email = strings.Split(credentials, ":")[0]
		got, err := testService.VerifyEmail(cred)

		assert.Nil(t, err)
		assert.Equal(t, "verified", got)
	})
	t.Run("Verfiy User try to retry", func(t *testing.T) {
		cred := db.Credentials{}

		credentials, _ := trestCommon.Decrypt("86c1e48f56bd8231f1be1f3946b7e7357991ba323f813957b68af0e2c1fd4df4a94349eea31f30e781343c9ddb6446b1c0b3c0879e5ad4acbc2c6e6a00cdb2705a")
		cred.VerificationCode = strings.Split(credentials, ":")[1]
		profileDB.Status = "verified"
		mockRepo := new(MockRepository)
		mockRepo.On("FindOne").Return(profileDB, nil)
		mockRepo.On("UpdateOne").Return("updated successfully", nil)
		testService := db.NewSignUpService(mockRepo)
		cred.Email = strings.Split(credentials, ":")[0]
		_, err := testService.VerifyEmail(cred)

		assert.NotNil(t, err)
		assert.Equal(t, "user already verified", err.Error())
	})
	t.Run("Verfiy User doesn't exist", func(t *testing.T) {

		cred := db.Credentials{}

		_, err := testService.VerifyEmail(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "email missing", err.Error())

	})
}

func TestVerifyResetLink(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindOne").Return(entity.ProfileDB{}, errors.New("user doesn't exist"))
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService := db.NewSignUpService(mockRepo)
	t.Run("Verfiy User Wrong Email", func(t *testing.T) {

		cred := db.Credentials{}
		cred.Email = "support@wiredsoft.org"
		cred.PasswordResetCode = profileDB.PasswordResetCode
		_, _, err := testService.VerifyResetLink(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "user doesn't exist", err.Error())

	})
	mockRepo = new(MockRepository)
	mockRepo.On("FindOne").Return(profileDB, nil)
	mockRepo.On("UpdateOne").Return("updated successfully", nil)
	testService = db.NewSignUpService(mockRepo)
	t.Run("Verfiy User Wrong code", func(t *testing.T) {

		cred := db.Credentials{}
		cred.Email = profileDB.Email
		cred.PasswordResetCode = "sdmklmdkljslakj"
		_, _, err := testService.VerifyResetLink(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "unauthorized user", err.Error())

	})
	t.Run("Verfiy User Success", func(t *testing.T) {
		cred := db.Credentials{}
		credentials, err := trestCommon.Decrypt("86c1e48f56bd8231f1be1f3946b7e7357991ba323f813957b68af0e2c1fd4df4a94349eea31f30e781343c9ddb6446b1c0b3c0879e5ad4acbc2c6e6a00cdb2705a")
		assert.Nil(t, err)
		cred.PasswordResetCode = strings.Split(credentials, ":")[1]

		cred.Email = strings.Split(credentials, ":")[0]
		got, _, err := testService.VerifyResetLink(cred)

		assert.Nil(t, err)
		assert.Equal(t, "verified", got)
	})
	t.Run("Verfiy User doesn't exist", func(t *testing.T) {

		cred := db.Credentials{}

		_, _, err := testService.VerifyResetLink(cred)
		assert.NotNil(t, err)
		assert.Equal(t, "email missing", err.Error())

	})
}

// func TestUpdatePassword(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	mockRepo.On("FindOne").Return(profileDB, nil)
// 	mockRepo.On("UpdateOne").Return("updated successfully", nil)
// 	testService := db.NewSignUpService(mockRepo)
// 	t.Run("Password Updated Successfully", func(t *testing.T) {
// 		cred := db.Credentials{}
// 		cred.Email = profileDB.Email
// 		got, err := testService.UpdatePassword(cred)
// 		assert.Nil(t, err)
// 		want := "password updated successfully"
// 		assert.Equal(t, got, want)
// 	})
// 	t.Run("Send Email without email", func(t *testing.T) {
// 		cred := db.Credentials{}
// 		cred.Email = profileDB.Email
// 		_, err := testService.UpdatePassword(cred)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, err.Error(), "password missing")
// 	})
// 	t.Run("Send Email without valid email", func(t *testing.T) {
// 		cred := db.Credentials{}
// 		cred.Email = "jhiugjkgkj"
// 		_, err := testService.UpdatePassword(cred)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, err.Error(), "invalid email")
// 	})
// 	mockRepo = new(MockRepository)
// 	mockRepo.On("FindOne").Return(entity.ProfileDB{}, errors.New("user doesnot exist"))
// 	mockRepo.On("UpdateOne").Return("updated successfully", nil)
// 	testService = db.NewSignUpService(mockRepo)
// 	t.Run("User doesnot exist", func(t *testing.T) {
// 		cred := db.Credentials{}
// 		cred.Email = "test@test.com"
// 		_, err := testService.UpdatePassword(cred)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, err.Error(), "user doesnot exist")
// 	})
// 	mockRepo = new(MockRepository)
// 	mockRepo.On("FindOne").Return(profileDB, nil)
// 	mockRepo.On("UpdateOne").Return("", errors.New("error updating password reset code"))
// 	testService = db.NewSignUpService(mockRepo)
// 	t.Run("Error While updating Password reset code in db", func(t *testing.T) {
// 		cred := db.Credentials{}
// 		cred.Email = "test@test.org"
// 		_, err := testService.UpdatePassword(cred)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, err.Error(), "error updating password reset code")
// 	})
// }
