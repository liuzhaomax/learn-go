package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewMediator()
	mediator.Ted.Talk()
}
