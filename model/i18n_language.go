package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type I18NLanguage struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code      string             `json:"code,omitempty" bson:"code,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Visible   int                `json:"visible,omitempty" bson:"visible,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updatedAt,omitempty"`
}

// IsExists struct
func (m I18NLanguage) IsExists() (ok bool) {
	if !m.ID.IsZero() {
		ok = true
	}

	return
}
