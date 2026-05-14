package factory

type IBed interface {
	SetSize(size int)
}

type Bed struct {
	size  int
	marca string
}

func NewBed(marca string) IBed {
	return &Bed{
		size:  0,
		marca: marca,
	}
}

func (b *Bed) SetSize(size int) {
	b.size = size
}

func (b *Bed) getProduct(product string) (interface{}, error) {
	return nil, nil
}
