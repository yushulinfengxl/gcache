package gcache

import "github.com/yushulinfengxl/gcache/internal/cache"

func Set(key string, value string) {
	cache.Set(key, value)
}

func Get(key string) (val string, err bool) {
	return cache.Get(key)
}
