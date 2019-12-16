package main

import "fmt"

func main() {
	a := 10
	b := "goland"
	c := 0.20
	x := 42
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Println("inside", x)
}

// fmt.Println("inside", x)
