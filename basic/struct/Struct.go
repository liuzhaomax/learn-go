package main

import "fmt"

type Person struct {
	Request Request
}

type Request struct {
	Name string
}

func main() {
	aaa := Person{}
	fmt.Println(aaa.Request == Request{}) // true
}
