package main

type A struct {
	field string
}

func exec(a A) {
	a.field = "world"
}

func exec1(a *A) {
	a.field = "world"
}

//func main() {
//	a := A{
//		field: "hello",
//	}
//	exec(a)
//	fmt.Println(a.field)
//	exec1(&a)
//	fmt.Println(a.field)
//}
