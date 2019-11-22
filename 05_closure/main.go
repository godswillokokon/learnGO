package main

import "fmt"

var x = 0

func incrementX() int {
	x++
	return x
}

func main() {
	fmt.Println(incrementX())
	fmt.Println(incrementX())

}
