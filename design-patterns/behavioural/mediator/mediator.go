package mediator

import "fmt"

type IMediator interface {
	Communicate(who string)
}

type IWildStallion interface {
	SetMediator(mediator IMediator)
}

type Bill struct {
	mediator IMediator
}

func (b *Bill) SetMediator(mediator IMediator) {
	b.mediator = mediator
}

func (b *Bill) Respond() {
	fmt.Println("Bill: what?")
	b.mediator.Communicate("Bill")
}

type Ted struct {
	mediator IMediator
}

func (t *Ted) SetMediator(mediator IMediator) {
	t.mediator = mediator
}

func (t *Ted) Respond() {
	fmt.Println("Ted: how much?")
}

func (t *Ted) Talk() {
	fmt.Println("Ted: Bill?")
	t.mediator.Communicate("Ted")
}

type ConcreteMediator struct {
	Bill
	Ted
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func (c *ConcreteMediator) Communicate(who string) {
	if who == "Ted" {
		c.Bill.Respond()
	} else {
		c.Ted.Respond()
	}
}
