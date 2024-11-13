package datastore

import "sync"

type DSEntry map[string]string

type Datastore struct {
	DSEntry
	Mutex *sync.RWMutex
}

func New() *Datastore {
	return &Datastore{
		DSEntry: make(DSEntry),
		Mutex:   new(sync.RWMutex),
	}
}

func (ds *Datastore) Set(key string, value string) {
	// 锁
	defer ds.Mutex.Unlock()
	ds.Mutex.Lock()
	ds.DSEntry[key] = value
}

func (ds *Datastore) Get(key string) string {
	defer ds.Mutex.RUnlock()
	ds.Mutex.RLock()
	return ds.DSEntry[key]
}
