package facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetCarBody() {
	fmt.Println("CarBody - SetCarBody")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}
