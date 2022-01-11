package share_repository

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

func NewSeeFewerStoriesRepository(collectionName string) SeeFewerStoriesRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document entity.SeeFewerStoriesDB) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert see fewer stories",
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
			"update see fewer stories",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("see fewer stories not found(404)")
		trestCommon.ECLog3(
			"update see fewer stories",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.SeeFewerStoriesDB, error) {
	var seefewerStories entity.SeeFewerStoriesDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&seefewerStories)
	if err != nil {
		trestCommon.ECLog3(
			"Find seefewerStories",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return seefewerStories, err
	}
	return seefewerStories, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.SeeFewerStoriesDB, error) {
	var seeFewerStories []entity.SeeFewerStoriesDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find seefewerStories",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var seeFewerStorie entity.SeeFewerStoriesDB
		if err = cursor.Decode(&seeFewerStorie); err != nil {
			trestCommon.ECLog3(
				"Find seefewerStories",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return seeFewerStories, nil
		}
		seeFewerStories = append(seeFewerStories, seeFewerStorie)
	}
	return seeFewerStories, nil
}
