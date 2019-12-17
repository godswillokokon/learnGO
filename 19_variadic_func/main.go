package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(24, 6, 284, 444, 999000, 836, 74, 8, 22202)
	fmt.Println(greatest)
}
