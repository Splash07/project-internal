package code

import (
	"context"
	"gitlab.com/Splash07/project-internal/model"
	"time"
)

// Create func
func (r Repo) Create(data model.I18NCode) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	// override data
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now
	data.Visible = 1
	//
		_, err = collection.InsertOne(context.Background(), data)
	return
}