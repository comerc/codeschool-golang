package main

import "fmt"

func main() {
	// array
	var langs [3]string
	langs[0] = "Go"
	langs[1] = "Ruby"
	langs[2] = "JavaScript"
	fmt.Println(langs)

	// slice
	var countries []string
	countries = append(countries, "Canada")
	countries = append(countries, "USA")
	countries = append(countries, "Mexico")
	fmt.Println(countries)

	// slice literals
	domains := []string{"COM", "NET", "ORG"}
	fmt.Println(domains[0])
	fmt.Println(domains[1])
	fmt.Println(domains[2])

	pets := getPets()
	for _, element := range pets {
		// fmt.Println(pets[i])
		fmt.Println(element)
	}
}

func getPets() []string {
	return []string{"cat", "dog"}
}
