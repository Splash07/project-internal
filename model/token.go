package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Token struct
type Token struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    int                `json:"userID,omitempty" bson:"userID,omitempty"`
	DeviceID  string             `json:"deviceID,omitempty" bson:"deviceID,omitempty"`
	VerifyKey string             `json:"verifyKey,omitempty" bson:"verifyKey,omitempty"`
	New       bool               `json:"new,omitempty" bson:"new,omitempty"`
	Token     string             `json:"token,omitempty" bson:"token,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	FcmID     string             `json:"fcmID,omitempty" bson:"fcmID,omitempty"`
}

// IsExists struct
func (m Token) IsExists() (ok bool) {
	if !m.ID.IsZero() {
		ok = true
	}

	return
}
