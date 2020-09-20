package code

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAll func
func (r Repo) GetAll(skip int, limit int) (results []model.I18NCode, total int64, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	option := options.Find()
	var cursor *mongo.Cursor
	// query := bson.M{"code": primitive.Regex{Pattern: code, Options: "im"}}
	total, err = collection.CountDocuments(context.Background(), bson.M{"visible": 1})
	cursor, err = collection.Find(
		context.Background(), bson.M{"visible": 1},
		option.SetLimit(int64(limit)),
		option.SetSkip(int64(skip)),
		option.SetSort(map[string]int{"createdAt": -1}))

	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			result := model.I18NCode{}
			err = cursor.Decode(&result)
			if err != nil {
				return nil, total, err
			}

			results = append(results, result)
		}

		// Close the cursor once finished
		_ = cursor.Close(context.TODO())
	}
	if err != nil {
		return
	}
	return
}
