package main

import "fmt"

const age = 70

func main() {
	// cant be changed
	const name string = "moeed"

	const age = 30

	fmt.Println("age>>>>>>>", age)

	const (
		port = 8000
		host = "localhost"
	)
	fmt.Println(port, host)

}
