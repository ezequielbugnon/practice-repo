package bitmanipulation

func IsEven(compared int) bool {
	return compared&1 == 0
}

type interruptor int

func NewInterruptor() interruptor {
	var i interruptor
	i = 1 << 0

	return i
}

func (i *interruptor) change() {
	*i ^= (1 << 0)
}
