package unregistered_user_repository

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

func NewUnRegisteredUserRepository(collectionName string) UnRegisteredUserRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document entity.UnRegisteredUsersDB) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert unregistereduser",
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
			"update unregistereduser",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("unregistereduser not found(404)")
		trestCommon.ECLog3(
			"update unregistereduser",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.UnRegisteredUsersDB, error) {
	var unregistereduser entity.UnRegisteredUsersDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&unregistereduser)
	if err != nil {
		trestCommon.ECLog3(
			"Find unregistereduser",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return unregistereduser, err
	}
	return unregistereduser, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.UnRegisteredUsersDB, error) {
	var unregisteredusers []entity.UnRegisteredUsersDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find unregistereduser",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var unregistereduser entity.UnRegisteredUsersDB
		if err = cursor.Decode(&unregistereduser); err != nil {
			trestCommon.ECLog3(
				"Find unregistereduser",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return unregisteredusers, nil
		}
		unregisteredusers = append(unregisteredusers, unregistereduser)
	}
	return unregisteredusers, nil
}
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete unregistereduser",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("unregistereduser not found(404)")
		trestCommon.ECLog3(
			"delete unregistereduser",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.UnRegisteredUsersDB, error) {
	var unregisteredusers []entity.UnRegisteredUsersDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find unregisteredusers",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var unregistereduser entity.UnRegisteredUsersDB
		if err = cursor.Decode(&unregistereduser); err != nil {
			trestCommon.ECLog3(
				"Find unregisteredusers",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return unregisteredusers, nil
		}
		unregisteredusers = append(unregisteredusers, unregistereduser)
	}
	return unregisteredusers, nil
}
