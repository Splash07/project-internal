package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type I18NCode struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code      string             `json:"code,omitempty" bson:"code,omitempty"`
	Languages []struct {
		LanguageID string `json:"language_id" bson:"languageId"`
		Text       string `json:"text" bson:"text"`
	} `json:"languages" bson:"languages"`
	Visible   int       `json:"visible,omitempty" bson:"visible"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedAt,omitempty"`
}

type I18nCodeTotal struct {
	Data  []I18NCode `json:"data"`
	Total int        `json:"total"`
}

// IsExists struct
func (m I18NCode) IsExists() (ok bool) {
	if !m.ID.IsZero() {
		ok = true
	}

	return
}
