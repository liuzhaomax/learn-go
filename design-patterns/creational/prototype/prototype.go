package prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Name() string {
	return c.name
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: c.name}
}
