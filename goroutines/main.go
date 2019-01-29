package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	names := []string{"Phil", "Noodles", "Barbaro", "Name4", "Name5", "Name6", "Name7", "Name8"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(name string, wg *sync.WaitGroup) {
	result := 0.0
	count := 100000000
	for index := 0; index < count; index++ {
		result += math.Pi * math.Sin(float64(len(name)))
	}
	fmt.Println("Name: ", name)
	wg.Done()
}
