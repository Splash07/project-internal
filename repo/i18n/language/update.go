package language

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-internal/model"
	"time"
)

// Update func
func (r Repo) Update(data model.I18NCode) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	// override data
	data.UpdatedAt = time.Now()
	//
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id" : data.ID},  bson.M{"$set": data})
	return
}
