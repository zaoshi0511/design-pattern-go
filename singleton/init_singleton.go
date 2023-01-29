package singleton

type InitSingleton struct{}

func init() {
	InitSingletonInstance = new(InitSingleton)
}

func GetInitSingletonInstance() *InitSingleton {
	return InitSingletonInstance
}

var InitSingletonInstance *InitSingleton
