package model

import (
	"time"
)

//System struct
type System struct {
	SystemID    int       `json:"system_id,omitempty" bson:"_id,omitempty"`
	Command     string    `json:"command,omitempty" bson:"command,omitempty"`
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Visible     bool      `json:"visible" bson:"visible"`
	VersionNo   string    `json:"version_no,omitempty" bson:"version_no,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy   int       `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy   int       `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	Description string    `json:"description" bson:"description"`
}

// IsExists struct
func (m System) IsExists() (ok bool) {
	if m.SystemID != 0 {
		ok = true
	}

	return
}
