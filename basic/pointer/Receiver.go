package pointer

import "fmt"

type Person struct {
	Name string
	Age  int
	Class
}

type Class struct {
	Name string
}

func (p *Person) SayName() {
	fmt.Println(p.Name)
}

func (p *Person) SayClassName() {
	fmt.Println(p.Class.Name)
}
