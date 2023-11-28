package memento

import (
	"fmt"
	"testing"
)

func TestNumber_ReinstateMemento(t *testing.T) {
	n := NewNumber(5)
	n.Double()
	n.Double()
	memento := n.CreateMemento()
	n.Half()
	n.ReinstateMemento(memento)
	fmt.Println(n.Value())
}
