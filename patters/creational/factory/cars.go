package factory

type ICar interface {
	setTypeMotor()
	isElectric() bool
	color()
}

type Car struct {
	name  string
	marca string
}

func NewCar(name, marca string) ICar {
	return &Car{
		name:  name,
		marca: marca,
	}
}

func (c *Car) setTypeMotor() {}
func (c *Car) isElectric() bool {
	return true
}
func (c *Car) color() {}
func (c *Car) getProduct(product string) (interface{}, error)
