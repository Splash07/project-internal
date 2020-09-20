package code

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-internal/model"
)

// GetOne func
func (r Repo) GetOne(id int) (result model.I18NCode, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	err = collection.FindOne(context.Background(), bson.M{"_id": id, "visible" : 1}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}
