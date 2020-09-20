package model

type Counter struct {
	ID       string `xorm:"_id" bson:"_id,omitempty"`
	Sequence int    `xorm:"seq" bson:"seq,omitempty"`
}
