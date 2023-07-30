# Go语言设计模式

此为个人学习设计模式时的笔记

## 一、创建型设计模式

### 1.单例模式

_注：以下几种单例模式的实现均为并发安全的。_

**懒加载式单例模式**

将实例的初始化延迟到获取时，并且使用Go语言内置的单机互斥锁`sync.Mutex`保证并发安全。优点是需要使用时才进行初始化，节省了内存；缺点是很慢，因为每次调用该函数/方法时都会进行加锁与解锁操作。

```go
package lazyload

import (
	"sync"
)

type singleton struct{}

var instance *singleton

var mutex sync.Mutex

func GetInstance() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
```

**双重检查式单例模式**

双重检查模式是懒加载模式的优化。通过两次对实例是否已初始化的判断，减少了加锁的可能性，从而提升了性能。

```go
package doublecheck

import "sync"

type singleton struct{}

var instance *singleton

var mutex sync.Mutex

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
```

**饥饿式加载单例模式**

饥饿式加载单例模式指在创建实例时便进行初始化，因此不存在并发安全的问题。在Java中通常通过类加载机制实现，而在Go语言中，是通过init函数在导入该包时便完成初始化操作。这种单例模式的优点是不存在并发安全问题，也不会因为重复调用初始化的方法造成资源的浪费；缺点则是实例自一开始便驻留在内存中，如果没有使用则会造成内存浪费。

```go
package hungryload

type singleton struct{}

var instance *singleton

func init() {
	if instance == nil {
		instance = &singleton{}
	}
}

func GetInstance() *singleton {
	return instance
}
```

**使用sync.Once实现单例模式**

`sync.Once`是Go语言SDK中所提供的一种结构体对象，其拥有的`Do`方法接受一个`Function Value`参数，可以保证传入的`Function Value`仅被执行一次。  
例子如下：显然这里的`Function Value`是一个闭包对象的地址而非二级函数指针，因为其捕获了一个外界的全局变量instance。

```go
package synconce

import "sync"

type singleton struct{}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
```

### 工厂方法模式