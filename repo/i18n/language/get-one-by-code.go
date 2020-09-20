package language

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-internal/model"
)

// GetOne func
func (r Repo) GetOneByCode(code string) (result model.I18NLanguage, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	err = collection.FindOne(context.Background(), bson.M{"code": code}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

