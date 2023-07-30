package FactoryMethodPattern

type Factory interface {
	FactoryMethod(owner string) Product
}
