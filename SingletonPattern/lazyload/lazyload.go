package lazyload

import (
	"sync"
)

type singleton struct{}

var instance *singleton

// 声明锁对象，保证协程并发安全
var mutex sync.Mutex

// GetInstance 懒加载式单例模式
// 使用互斥锁保证了并发安全，但每次调用都需要加锁和解锁，性能较低
func GetInstance() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
