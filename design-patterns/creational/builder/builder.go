package builder

type Concrete struct {
	name  string
	built bool
}

type ConcreteBuilder struct {
	name  string
	built bool
}

func NewConcreteBuilder(name string) *ConcreteBuilder {
	return &ConcreteBuilder{name, false}
}

func (b *ConcreteBuilder) setBuilt(built bool) *ConcreteBuilder {
	b.built = built
	return b
}

func (b *ConcreteBuilder) Build() *Concrete {
	// 这里可以加校验
	concrete := &Concrete{
		name:  b.name,
		built: b.built,
	}
	return concrete
}

type Product struct {
	Name  string
	Built bool
}

func (b *Concrete) GetResult() Product {
	return Product{b.name, b.built}
}
