package model

import "time"

type CommandPath struct {
	Command   string    `json:"command,omitempty" bson:"command,omitempty"`
	Path      string    `json:"path,omitempty" bson:"path,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
