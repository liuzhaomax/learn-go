package memento

type Memento struct {
	state int
}

func NewMemento(value int) *Memento {
	return &Memento{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value: value}
}

func (n *Number) Double() {
	n.value = n.value * 2
}

func (n *Number) Half() {
	n.value /= 2
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}

func (n *Number) ReinstateMemento(memento *Memento) {
	n.value = memento.state
}
