package main

import "fmt"

func main() {
	switch "Apple" {
	case "Orange":
		fmt.Println("Oranges don't have orange color")
	case "Grapes":
		fmt.Println("Grapes are sour")
	case "Apple":
		fmt.Println("An Apple a day keeps you away from a doctor")
		fallthrough
	case "Apples":
		fmt.Println("I Love green apples")
	default:
		fmt.Println("No fruit")
	}
}
