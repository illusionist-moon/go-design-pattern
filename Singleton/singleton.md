# 单例模式

单例模式是所有设计模式中最为简单且好理解的，但即使简单，也有着许多需要注意的细节。
<br>
在实际的项目中，有一些资源或者实例我们需要只创建一份，如果重复创建，则会造成不必要的性能损耗甚至导致严重的错误。例如各种池（如：协程池、连接池...），以及与一些网络构件的连接（如：数据库）等。
<br>

单例模式从大体上分为饿汉式与懒汉式，下面将逐个介绍。

## 1、饿汉式单例模式

饿汉式单例模式指在创建实例时便进行初始化，因此不存在并发安全的问题。在Java中通常通过类初始化时的类加载机制实现，而在Go语言中，是通过init函数在导入该包时便完成初始化操作。
<br>
这种单例模式的优点是不存在并发安全问题，也不会因为重复调用初始化的方法造成性能损耗；缺点则是实例自一开始便驻留在内存中，如果没有使用则会造成内存浪费。

```go
package hungryload

type singleton struct{}

func (s *singleton) Work() {}

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

上述代码便是饿汉式单例模式的一个简单化的实现，看似好像已经OK了，但其实还存在需要讨论的问题。
<br>
我们可以看到，instance的类型显然是未导出的，这在Java中通常是private的，即我们不希望接收方拿到instance后对其进行一些不合法的修改。但事实上，这种写法会导致位于不同包的接收方拿到instance后根本无法做任何事情，无论是作为函数或方法的参数还是调用instance的方法。因此我们需要进行一些改变。
<br>
在计算机领域中，没有什么是多加一层不能解决的。我们很容易能想到利用接口来进行解耦。代码如下：

```go
package hungryload

type IInstance interface {
	Work()
}

type singleton struct{}

func (s *singleton) Work() {}

var instance *singleton

func init() {
	if instance == nil {
		instance = &singleton{}
	}
}

func GetInstance() IInstance {
	return instance
}

```

## 懒汉式单例模式

懒汉式单例模式将实例的初始化延迟到真正需要时，并且使用Go语言内置的单机互斥锁`sync.Mutex`保证并发安全。
<br>
优点是需要使用时才进行初始化，节省了内存；缺点是很慢，因为每次调用该函数/方法时都会进行加锁与解锁操作。

```go
package lazyload

import (
	"sync"
)

type IInstance interface {
	Work()
}

type singleton struct{}

func (s *singleton) Work() {}

var instance *singleton

// 声明锁对象，保证协程并发安全
var mutex sync.Mutex

// GetInstance 懒汉式单例模式
// 使用互斥锁保证了并发安全，但每次调用都需要加锁和解锁，性能较低
func GetInstance() IInstance {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
```

**双重检查式单例模式**

接下来我们考虑这样一个问题：我们之所以使用锁，是为了在实例还没有初始化时避免多个协程并发调用初始化的函数，造成重复创建。而一旦这个实例已经被初始化后，外部再调用这个方法事实上是没有必要加锁的，因为在判断时就会跳转至return语句。我们都知道，互斥锁的性能是很差的，因此我们应该尽可能去减少使用锁的可能。
<br>
双重检查模式是便是上述懒加载模式的优化。通过两次对实例是否已初始化的判断，减少了加锁的可能性。因为一条判断语句与加锁带来的性能消耗相比完全可以忽略不计，从而提升了性能。

```go
package doublecheck

import "sync"

type singleton struct{}

var instance *singleton

var mutex sync.Mutex

func GetInstance() IInstance {
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

**使用sync.Once实现单例模式**

`sync.Once`是Go语言SDK中所提供的一种结构体对象，其拥有的`Do`方法接受一个`Function Value`参数，可以保证传入的`Function Value`仅被执行一次。
<br>
事实上，`sync.Once`就是Go语言标准库对双重检测模式的一种封装。

例子如下：显然这里的`Function Value`是一个闭包对象的地址而非二级函数指针，因为其捕获了一个外界的全局变量instance。

```go
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
```