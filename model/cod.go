package model

import "time"

type Debt struct {
	ID       int `json:"_id" bson:"_id,omitempty"`
	ClientID int `json:"client_id" bson:"client_id,omitempty" structs:"client_id"`
	Amount   int `json:"amount" bson:"amount,omitempty"`

	//
	UpdatedIP       string    `json:"updated_ip" bson:"updated_ip,omitempty"`
	UpdatedEmployee int       `json:"updated_employee" bson:"updated_employee,omitempty"`
	UpdatedClient   int       `json:"updated_client" bson:"updated_client,omitempty"`
	UpdatedSource   string    `json:"updated_source" bson:"updated_source,omitempty"`
	UpdatedDate     time.Time `json:"updated_date" bson:"updated_date,omitempty"`
}
