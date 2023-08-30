package synconce

import "sync"

type IInstance interface {
	Work()
}

type singleton struct{}

func (s *singleton) Work() {}

var instance *singleton

// 使用Go标准库提供的sync.Once，做到使函数仅加载一次
var once sync.Once

func GetInstance() IInstance {
	// once.Do 接受一个 Function Value 参数
	// 这里的 Function Value 显然是一个闭包对象的地址而不是二级函数指针
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
