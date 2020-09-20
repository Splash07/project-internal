package code

import (
	"sync"

	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-internal/repo"
)

// Repo type
type Repo repo.Mongo

var (
	instance *Repo
	once     sync.Once
)

// New ..
func New() *Repo {
	once.Do(func() {
		instance = &Repo{
			Session:    config.GetConfig().MongoV2.Get("internal"),
			Collection: "i18n_code",
		}
	})
	return instance
}
