package doublecheck

import "sync"

type singleton struct{}

var instance *singleton

// 声明锁对象，保证协程并发安全
var mutex sync.Mutex

// GetInstance 双重检查单例模式
// 是懒加载单例模式的优化，使用两次判断减少了使用锁的可能性
func GetInstance() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = &singleton{}
		}
		mutex.Unlock()
	}
	return instance
}
