package main

import "fmt"

func main() {
	x := 10 % 2
	fmt.Println(x)
	if x >= 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
