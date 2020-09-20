package model

import (
	"time"
)

// Token struct
type User struct {
	UserID    int       `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Status    string    `json:"status,omitempty" bson:"status,omitempty"`
	Level     int       `json:"level,omitempty" bson:"level,omitempty"`
	VersionNo string    `json:"version_no,omitempty" bson:"version_no,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt,omitempty"`
	CreatedBy int       `json:"created_by,omitempty" bson:"createdBy,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedAt,omitempty"`
	UpdatedBy int       `json:"updated_by,omitempty" bson:"updatedBy,omitempty"`
}

// IsExists struct
func (m User) IsExists() (ok bool) {
	if m.UserID > 0 {
		ok = true
	}

	return
}
