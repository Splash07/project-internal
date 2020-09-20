package language

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAll func
func (r Repo) GetAll(id string, limit int) (results []model.I18NLanguage, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	option := options.Find()
	var cursor *mongo.Cursor
	ObjectID, _ := primitive.ObjectIDFromHex(id)
	cursor, err = collection.Find(context.Background(), bson.M{"_id": bson.M{"$gte": ObjectID}, "visible": 1}, option.SetLimit(int64(limit)))

	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			result := model.I18NLanguage{}
			err = cursor.Decode(&result)
			if err != nil {
				return nil, err
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
