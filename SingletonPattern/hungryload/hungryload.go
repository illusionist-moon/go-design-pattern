package hungryload

type singleton struct{}

var instance *singleton

// 饥饿式单例
// 通过init函数，在导入包时便创建对象
// 并发安全，唯一的缺点是创建的对象无论是否被使用，都会持续存储在内存中
func init() {
	if instance == nil {
		instance = &singleton{}
	}
}

func GetInstance() *singleton {
	return instance
}
