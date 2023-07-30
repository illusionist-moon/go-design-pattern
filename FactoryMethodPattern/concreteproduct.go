package FactoryMethodPattern

import "fmt"

// 用于检测 ConcreteProduct 类是否实现了 Product 接口
var _ Product = &ConcreteProduct{}

type ConcreteProduct struct {
	owner string
}

func (cp *ConcreteProduct) Use() {
	fmt.Printf("owner: %s\n", cp.owner)
}
