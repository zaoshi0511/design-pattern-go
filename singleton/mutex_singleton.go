package singleton

import "sync"

type MutexSingleton struct{}

func GetMutexSingletonInstance() *MutexSingleton {
	if nil == MutexSingletonInstance {
		mu.Lock()
		defer mu.Unlock()
		if nil == MutexSingletonInstance {
			MutexSingletonInstance = &MutexSingleton{}
		}
	}
	return MutexSingletonInstance
}

var (
	MutexSingletonInstance *MutexSingleton
	mu                     sync.Mutex
)
