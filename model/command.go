package model

import "time"

type Command struct {
	CommandID int       `json:"_id,omitempty" bson:"_id,omitempty"`
	ModuleID  int       `json:"module_id,omitempty" bson:"module_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Visible   bool      `json:"visible" bson:"visible"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy int       `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy int       `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
}

// IsExists struct
func (m Command) IsExists() (ok bool) {
	if m.CommandID > 0 {
		ok = true
	}
	return
}
