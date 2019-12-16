package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func marlians() {
	fmt.Print(" Marlians")
}
func main() {
	defer marlians()
	hello()
}
