package singleton

import "sync"

type SyncSingleton struct{}

func GetSyncSingletonInstanceInstance() *SyncSingleton {
	once.Do(func() {
		SyncSingletonInstance = &SyncSingleton{}
	})
	return SyncSingletonInstance
}

var (
	SyncSingletonInstance *SyncSingleton
	once                  sync.Once
)
