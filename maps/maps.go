package main

import "fmt"

// map > hash, object , dictionary
func main() {

	// creating map

	// m := make(map[string]string)

	// //setting a element\
	// m["name"] = "moeed"
	// m["area"] = "backend"

	// //geeting element
	// fmt.Println("this is mapp", m["name"], m["area"])
	// fmt.Println("empty", m["ok"])

	mymap := map[string]int{"age": 20}

	fmt.Println("this is mymap", mymap["age"])

	a, ok := mymap["ages"]
	fmt.Println(ok)
	fmt.Println("this>>>>>", a)

	if ok {
		fmt.Println("yes all ok")
	} else {
		fmt.Println("not ok")
	}

}
