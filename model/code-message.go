package model

import "time"

// CodeMessage ..
type CodeMessage struct {
	ID       string    `json:"_id"`
	Code     string    `json:"code"`
	Text     string    `json:"text"`
	Visible  int       `json:"visible"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

// IsExists ..
func (m CodeMessage) IsExists() (ok bool) {
	if m.ID != "" {
		ok = true
	}
	return
}
