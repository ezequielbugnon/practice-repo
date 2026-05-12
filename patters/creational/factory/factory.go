package factory

type IFactory interface {
	getProduct(product string) (interface{}, error)
}

func Factory(product, witch string) interface{} {
	if product == "car" {
		return NewCar("", "")
	}

	return nil
}
