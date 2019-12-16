package main

import "fmt"

func main() {

	var year int
	fmt.Print("What year do you want to check ? : ")
	fmt.Scan(&year)
	results := year % 4
	if results == 0 {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}

}
