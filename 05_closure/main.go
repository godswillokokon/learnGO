package main

import "fmt"

var x = 0

func wrapper() func() int {
	z := 50
	return func() int {
		x++
		return z
	}
}

func incrementX() int {
	x++
	return x
}

func main() {
	incrementZ := wrapper()
	fmt.Println(incrementZ())
	fmt.Println(incrementZ())

	fmt.Println(incrementX())
	fmt.Println(incrementX())

	y := 10
	incrementY := func() int {
		y++
		return y
	}
	fmt.Println(incrementY())
	fmt.Println(incrementY())

}
