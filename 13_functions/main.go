package main

import "fmt"

func main() {
	s := greet("Bassey ", "Okon")
	fmt.Println(s)
	fmt.Println(greets(" Godswill ", " Okokon "))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func greets(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
