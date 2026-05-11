package builder

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildRag() *Rag {
	d.builder.setAgent()
	d.builder.setFramework()
	d.builder.setLlm()
	d.builder.setVectorDB()
	return d.builder.getRag()
}
