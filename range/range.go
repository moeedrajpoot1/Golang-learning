package main

import "fmt"

func main() {

	//. iterating over data structue
	// nums := []int{1, 2, 3, 4, 5}

	// for i := 0; i <= len(nums); i++ {

	// 	fmt.Println("this is helllo ", i)
	// }
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println("num>>>>>>>", num)
	// 	fmt.Println("this is index>>>>", i)
	// 	sum += num
	// }
	// fmt.Println("this is sum>>>>>>>>>", sum)

	m := map[string]string{"name": "moeed", "lname": "testing"}

	for i, name := range m {
		fmt.Println("this is name>>>>>>>>>>>>.", name)
		fmt.Println("this is index>>>>>>>>>>>", i)
	}

}
