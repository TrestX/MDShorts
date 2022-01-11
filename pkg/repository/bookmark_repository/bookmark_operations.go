package bookmark_repository

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

func NewBookmarkRepository(collectionName string) BookMarkRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

func (r *repo) InsertOne(document entity.BookmarkDB) (string, error) {
	user, err := trestCommon.InsertOne(document, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"insert bookmark",
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
			"update bookmark",
			err,
			logrus.Fields{
				"filter":          filter,
				"update":          update,
				"collection name": r.CollectionName,
			})

		return "", err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		err = errors.New("bookmark not found(404)")
		trestCommon.ECLog3(
			"update bookmark",
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

func (r *repo) FindOne(filter, projection bson.M) (entity.BookmarkDB, error) {
	var bookmark entity.BookmarkDB
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&bookmark)
	if err != nil {
		trestCommon.ECLog3(
			"Find bookmark",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return bookmark, err
	}
	return bookmark, err
}

func (r *repo) Find(filter, projection bson.M) ([]entity.BookmarkDB, error) {
	var bookmarks []entity.BookmarkDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 200, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find profiles",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var bookmark entity.BookmarkDB
		if err = cursor.Decode(&bookmark); err != nil {
			trestCommon.ECLog3(
				"Find profiles",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return bookmarks, nil
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}

func (r *repo) FindNewsIds(filter, projection bson.M) ([]entity.BMNId, error) {
	var bookmarks []entity.BMNId
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 200, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find newsids",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var bookmark entity.BMNId
		if err = cursor.Decode(&bookmark); err != nil {
			trestCommon.ECLog3(
				"Find newsids",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return bookmarks, nil
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}

func (r *repo) DeleteOne(filter bson.M) error {
	deleteResult, err := trestCommon.DeleteOne(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"delete bookmark",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	if deleteResult.DeletedCount == 0 {
		err = errors.New("bookmark not found(404)")
		trestCommon.ECLog3(
			"delete bookmark",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return err
	}
	return nil
}

func (r *repo) FindWithIDs(filter, projection bson.M) ([]entity.BookmarkDB, error) {
	var bookmarks []entity.BookmarkDB
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 200, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find bookmarks",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var bookmark entity.BookmarkDB
		if err = cursor.Decode(&bookmark); err != nil {
			trestCommon.ECLog3(
				"Find bookmarks",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return bookmarks, nil
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}
func (r *repo) FindWithAggregate(filter primitive.A) ([]entity.BookMarkADB, error) {
	var bookmarks []entity.BookMarkADB
	cursor, err := trestCommon.Aggregate(filter, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find bookmarks",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var bookmark entity.BookMarkADB
		if err = cursor.Decode(&bookmark); err != nil {
			trestCommon.ECLog3(
				"Find bookmarks",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return bookmarks, nil
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}
