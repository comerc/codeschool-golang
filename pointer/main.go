package main

import "fmt"

func main() {
	gopher1 := &gopher{name: "Phil", age: 30}
	validateAge(gopher1)
	fmt.Println(*gopher1)
}

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}
