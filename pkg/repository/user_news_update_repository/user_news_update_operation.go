package user_news_update_repository

import (
	"MdShorts/pkg/entity"
	"context"
	"errors"

	"github.com/aekam27/trestCommon"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repo struct {
	CollectionName string
}

func NewUserNewsUpdateRepository(collectionName string) UserNewsUpdateRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document entity.UserNewsUpdateDB) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert user news update",
			err,
			logrus.Fields{
				"document":        document,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	userid := user.InsertedID.(primitive.ObjectID).Hex()
	return userid, nil
}

func (r *repo) UpdateOne(filter, update bson.M) (string, error) {
	result, err := trestCommon.UpdateOne(filter, update, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"update user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("user news update not found(404)")
		trestCommon.ECLog3(
			"update user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})
		return "", err
	}
	return "updated successfully", nil
}

func (r *repo) FindOne(filter, projection bson.M) (entity.UserNewsUpdateDB, error) {
	var userNewsUpdate entity.UserNewsUpdateDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&userNewsUpdate)
	if err != nil {
		trestCommon.ECLog3(
			"Find user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return userNewsUpdate, err
	}
	return userNewsUpdate, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.UserNewsUpdateDB, error) {
	var userNewsUpdates []entity.UserNewsUpdateDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var userNewsUpdate entity.UserNewsUpdateDB
		if err = cursor.Decode(&userNewsUpdate); err != nil {
			trestCommon.ECLog3(
				"Find user news update",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return userNewsUpdates, nil
		}
		userNewsUpdates = append(userNewsUpdates, userNewsUpdate)
	}
	return userNewsUpdates, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("user news update not found(404)")
		trestCommon.ECLog3(
			"delete user news update",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
