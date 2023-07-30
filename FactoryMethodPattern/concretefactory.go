package FactoryMethodPattern

// 用于检测ConcreteFactory类是否实现了Factory接口
var _ Factory = &ConcreteFactory{}

type ConcreteFactory struct {
}

func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{owner: owner}
	return p
}
