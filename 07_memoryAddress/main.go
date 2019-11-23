package main

import "fmt"

const metersToYards float64 = 1.09861

func main() {

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters - metersToYards
	fmt.Println(meters, "meters is", yards, "yards")

	// a := 22
	// b := 1
	// c := 20
	// fmt.Println("a value is ", a)
	// fmt.Println("a memory address is ", &a)
	// fmt.Printf("%d \n", &a)
	// fmt.Printf("%d \n", &b)
	// fmt.Printf("%d \n", &c)

}
