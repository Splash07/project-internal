package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserPermission struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PermissionID int                `json:"permission_id,omitempty" bson:"permission_id,omitempty"`
	UserID       int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Visible      bool               `json:"visible" bson:"visible"`
	Note         string             `json:"note,omitempty" bson:"note,omitempty"`
	VersionNo    string             `json:"version_no,omitempty" bson:"version_no,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy    int                `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy    int                `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
}

// IsExists struct
func (m UserPermission) IsExists() (ok bool) {
	if !m.ID.IsZero() {
		ok = true
	}
	return
}
