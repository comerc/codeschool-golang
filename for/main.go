package main

func main() {
	println("x")
	for x := 0; x < 5; x++ {
		println(x)
	}
	y := 0
	is := true
	println("y")
	for is {
		if y > 5 {
			is = false
		}
		y++
		println(y)
	}
	z := 0
	println("z")
	for {
		if z > 3 {
			break
		}
		z++
		println(z)
	}
}
