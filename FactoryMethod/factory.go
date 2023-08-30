package FactoryMethod

type Factory interface {
	FactoryMethod(owner string) Product
}
