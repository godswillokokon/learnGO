package main

import "fmt"

func main() {
	n := avg(19, 38, 33, 98, 39, 37, 99)
	fmt.Println(n)
}

func avg(nums ...float64) float64 {
	fmt.Println(nums)
	fmt.Printf("%T \n", nums)
	// this sets thr var total to 0
	var total float64
	for _, a := range nums {
		total += a
	}
	return total / float64(len(nums))
}
