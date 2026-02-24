package main

import "fmt"

// functions can return multiple values on go
// func add(a int, b int) int {

// 	return a + b

// }

// func getlanguages() (string, string, string) {

// 	return "Moeed", "Anees ", "Awais"

// }

func processit() func(a int) int {
	return func(a int) int {
		return 4
	}

}

func main() {

	fn := processit()

	fmt.Println("this is processs>>>>>>>>>>>>>", fn(6))

	// fmt.Println("this is. a result", add(2, 3))

	// lang1, lang2, lang3 := getlanguages()
	// fmt.Println("These are the get languages", lang1)
	// fmt.Println("this is lang2>>>>>>>>", lang2)
	// fmt.Println("this is lang3", lang3)
}
