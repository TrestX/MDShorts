package db

import (
	"MdShorts/pkg/entity"
	bookmark "MdShorts/pkg/repository/bookmark_repository"
	news_repository "MdShorts/pkg/repository/news_repository"
	"errors"
	"time"

	"github.com/aekam27/trestCommon"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo = bookmark.NewBookmarkRepository("bookmark")
)
var (
	newsrepo = news_repository.NewNewsRepository("news")
)

type bookmarkService struct{}

func NewBookmarkService(repository bookmark.BookMarkRepository) BookmarkService {
	repo = repository
	return &bookmarkService{}
}
func (*bookmarkService) AddBookmark(bookmark BookMark) (string, error) {
	var bookmarkEntity entity.BookmarkDB
	if bookmark.NewsId == "" {
		return "", errors.New("newsid missing")
	}
	if bookmark.UserId == "" {
		return "", errors.New("userid missing")
	}
	bD, err := checkByNIDUID(bookmark.UserId, bookmark.NewsId)
	if err == nil {
		if bD.Status == "Active" {
			return "", errors.New("bookmark already exist")
		} else {
			id, _ := primitive.ObjectIDFromHex(bD.ID.Hex())
			setParameters := bson.M{}
			setParameters["status"] = "Active"
			setParameters["updated_time"] = time.Now()
			filter := bson.M{"_id": id}
			set := bson.M{
				"$set": setParameters,
			}
			return repo.UpdateOne(filter, set)
		}
	}
	bookmarkEntity.ID = primitive.NewObjectID()
	bookmarkEntity.Status = "Active"
	bookmarkEntity.AddedTime = time.Now()
	bookmarkEntity.NewsId = bookmark.NewsId
	bookmarkEntity.UserId = bookmark.UserId
	return repo.InsertOne(bookmarkEntity)
}

func (*bookmarkService) UpdateBookmarkStatus(bookmark BookMark, bookmarkid string) (string, error) {
	if bookmarkid == "" {
		err := errors.New("bookmark id missing")
		trestCommon.ECLog2(
			"update bookmark",
			err,
		)
		return "", err
	}

	bD, err := checkByNIDUID(bookmark.UserId, bookmark.NewsId)
	if err != nil {
		return "", errors.New("invalid bookmark Id")
	}
	id, _ := primitive.ObjectIDFromHex(bD.ID.Hex())
	setParameters := bson.M{}
	if bookmark.Status != "" {
		setParameters["status"] = bookmark.Status
	}
	setParameters["updated_time"] = time.Now()
	filter := bson.M{"_id": id}
	set := bson.M{
		"$set": setParameters,
	}
	return repo.UpdateOne(filter, set)
}

func checkByBookmarkID(id primitive.ObjectID) (entity.BookmarkDB, error) {
	bookmark, err := repo.FindOne(bson.M{"_id": id}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Bookmark Details section",
			err,
		)
		return bookmark, err
	}
	return bookmark, nil
}

func checkByNIDUID(id string, nid string) (entity.BookmarkDB, error) {
	bookmark, err := repo.FindOne(bson.M{"user_id": id, "newsId": nid}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Bookmark Details section",
			err,
		)
		return bookmark, err
	}
	return bookmark, nil
}

func (*bookmarkService) GetBookmarks(limit, skip int, status, userid, newsid string) ([]entity.NewsDB, error) {
	// filter := bson.M{
	// 	"$match": bson.M{"status": "Active"},
	// }
	// if userid != "" {
	// 	filter = bson.M{
	// 		"$match": bson.M{
	// 			"user_id": userid,
	// 			"status":  "Active",
	// 		},
	// 	}
	// }

	// aggfilter := bson.A{
	// 	filter,
	// 	bson.M{
	// 		"$project": bson.M{
	// 			"newsOId": bson.M{
	// 				"$toObjectId": "$newsId",
	// 			},
	// 		},
	// 	},
	// 	bson.M{
	// 		"$lookup": bson.M{
	// 			"from":         "news",
	// 			"localField":   "newsOId",
	// 			"foreignField": "_id",
	// 			"as":           "bnews",
	// 		},
	// 	},
	// 	bson.M{
	// 		"$match": bson.M{
	// 			"bnews": bson.M{
	// 				"$ne": []interface{}{},
	// 			},
	// 		},
	// 	},
	// 	bson.M{
	// 		"$project": bson.M{
	// 			"newsOId": 1,
	// 			"_id":     1,
	// 			"newss":   bson.M{"$arrayElemAt": []interface{}{"$bnews", 0}},
	// 		},
	// 	},
	// }
	filter := bson.M{"status": "Active"}
	if userid != "" {
		filter["user_id"] = userid
	}
	bookmark, err := repo.Find(filter, bson.M{})
	if err != nil {
		return []entity.NewsDB{}, err
	}
	subFilter := bson.A{}
	for _, nid := range bookmark {
		id, _ := primitive.ObjectIDFromHex(nid.NewsId)
		subFilter = append(subFilter, bson.M{"_id": id})
	}
	filter1 := bson.M{"$or": subFilter}
	news, err := newsrepo.FindWithIDs(filter1, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get bookmark section",
			err,
		)
		return []entity.NewsDB{}, err
	}
	return news, nil
}

func (*bookmarkService) GetBookmarksIDs(limit, skip int, status, userid, newsid string) ([]entity.BMNId, error) {
	bookmark, err := repo.FindNewsIds(bson.M{"user_id": userid}, bson.M{})
	if err != nil {
		trestCommon.ECLog2(
			"Get Bookmark Details section",
			err,
		)
		return bookmark, err
	}
	return bookmark, nil
}
