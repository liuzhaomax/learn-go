package composition

import "fmt"

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value: value}
}

func (l *Leaf) Traverse() {
	fmt.Println(l.value)
}

type Composition struct {
	children []Component
}

func NewComposition() *Composition {
	return &Composition{children: make([]Component, 0)}
}

func (c *Composition) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composition) Traverse() {
	for i := range c.children {
		c.children[i].Traverse()
	}
}
