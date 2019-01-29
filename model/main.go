package model

import "fmt"

type gopher struct {
	name string
	age  int
}

type horse struct {
	name   string
	weight int
}

type jumper interface {
	Jump() string
}

func (h *horse) Jump() string {
	if h.weight > 2500 {
		return "I'm tooheavy, can't jump..."
	}
	return "I will jump, Neigh!!"
}

func (g *gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}

func main() {
	jumperList := GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}

func GetList() []jumper {
	gopher1 := &gopher{name: "Phil", age: 30}
	gopher2 := &gopher{name: "Noodles", age: 90}
	horse1 := &horse{name: "Barbaro", weight: 2000}
	list := []jumper{gopher1, gopher2, horse1}
	return list
}
