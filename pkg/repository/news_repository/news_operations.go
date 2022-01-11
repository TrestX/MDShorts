package news_repository

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

func NewNewsRepository(collectionName string) NewsRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document entity.NewsDB) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert news",
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
			"update news",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("news not found(404)")
		trestCommon.ECLog3(
			"update news",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.NewsDB, error) {
	var news entity.NewsDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&news)
	if err != nil {
		trestCommon.ECLog3(
			"Find news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return news, err
	}
	return news, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	cursor, err := trestCommon.Find(filter, projection, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var new entity.NewsDB
		if err = cursor.Decode(&new); err != nil {
			trestCommon.ECLog3(
				"Find news",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return news, nil
		}
		news = append(news, new)
	}
	return news, nil
}
func (r *repo) FindSort(filter, filter1, projection bson.M, limit, skip int) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	cursor, err := trestCommon.FindSort(filter, projection, filter1, limit, skip, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var new entity.NewsDB
		if err = cursor.Decode(&new); err != nil {
			trestCommon.ECLog3(
				"Find news",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return news, nil
		}
		news = append(news, new)
	}
	return news, nil
}
func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("news not found(404)")
		trestCommon.ECLog3(
			"delete news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}
func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.NewsDB, error) {
	var news []entity.NewsDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 200, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find news",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var new entity.NewsDB
		if err = cursor.Decode(&new); err != nil {
			trestCommon.ECLog3(
				"Find news",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return news, nil
		}
		news = append(news, new)
	}
	return news, nil
}
