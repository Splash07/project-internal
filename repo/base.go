package repo

import (
	"gitlab.ghn.vn/online/common/adapter/v2"
)

// Mongo struct
type Mongo struct {
	Session    *adapter.MongoV2
	Collection string
}

// MSSQL struct
type MSSQL struct {
	Session    *adapter.MSSQL
	Table      string
	PrimaryKey string
}
