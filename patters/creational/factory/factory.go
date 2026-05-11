package factory

type IFactory interface {
	getProduct(product string) (interface{}, error)
}
