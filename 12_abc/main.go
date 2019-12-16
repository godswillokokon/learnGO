package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "--A--")
		} else if i%3 == 0 {
			fmt.Println("--B--")
		} else if i%5 == 0 {
			fmt.Println(i, "--C--")
		} else {
			fmt.Println(i)
		}

	}
}
