package main

import "fmt"

func main() {
	a := 19

	fmt.Println(a)  //19
	fmt.Println(&a) //hexadecimal

	var b *int = &a
	fmt.Println(b)  //hexadecimal
	fmt.Println(*b) //19
}
