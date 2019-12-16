package main

import "net/http"
import "log"
import "io/ioutil"
import "fmt"

func main() {
	res, err := http.Get("https://og-movies.herokuapp.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
