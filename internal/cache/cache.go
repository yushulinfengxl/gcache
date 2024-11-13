package cache

import (
	"github.com/yushulinfengxl/ydb/internal/datastore"
	"github.com/yushulinfengxl/ydb/utility/singleton"
)

type Cache struct {
	Db *datastore.Datastore
}

var cacheLazySingleton *singleton.Lazy

func New() *Cache {
	ins := cacheLazySingleton.Instance(&Cache{
		Db: datastore.New(),
	})
	return (*ins).(*Cache)
}

func Set(key string, value string) {
	New().Db.Set(key, value)
}

func Get(key string) (val string, err bool) {
	val = New().Db.Get(key)
	return
}
