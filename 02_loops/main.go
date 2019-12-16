package main

import "fmt"

func main() {
		for index := 1; index < 50; index++ {
		if index % 2 ==1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

}
