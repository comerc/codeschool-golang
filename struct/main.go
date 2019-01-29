package main

import "fmt"

func main() {
	gopher1 := gopher{name: "Phil", age: 30}
	fmt.Println(gopher1.jump())
	gopher2 := gopher{name: "Noodles", age: 90}
	fmt.Println(gopher2.jump())
}

type gopher struct {
	name string
	age  int
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}
