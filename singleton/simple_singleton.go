package singleton

type SimpleSingleton struct{}

func GetSimpleSingletonInstance() *SimpleSingleton {
	if nil == SimpleSingletonInstance {
		SimpleSingletonInstance = &SimpleSingleton{}
	}
	return SimpleSingletonInstance
}

var SimpleSingletonInstance *SimpleSingleton
