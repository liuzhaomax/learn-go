package pointer

import "testing"

func TestPerson_SayClassName(t *testing.T) {
	person1 := &Person{
		Name:  "xiaoming",
		Age:   1,
		Class: Class{Name: "1ban"},
	}
	person1.SayClassName()
}
