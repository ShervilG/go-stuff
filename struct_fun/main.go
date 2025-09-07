package main

import "fmt"

type Child struct {
	name string
	age  int
}

func main() {
	firstChild := &Child{name: "Shervil", age: 27}
	fmt.Printf("%s:%v", firstChild.name, firstChild.age)
}
