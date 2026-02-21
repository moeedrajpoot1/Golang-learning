package main

import "fmt"

func main() {
	// only have for loop dont have other

	// while loop
	// i := 1

	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1

	// }

	// // infinte loop

	// for {
	// 	fmt.Println("start")
	// }

	// // classic for loop

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("this is i", i)
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println("testing")
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 100; i++ {

	// 	if i%2 == 0 {
	// 		fmt.Println("even number", i)
	// 	}
	// }

	// // sum
	// sum := 0

	// for i := 1; i <= 10; i++ {

	// 	sum += i
	// }
	// fmt.Println("this is total sum", sum)
	// // table

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("2 x ", i, "=", i*2)
	// }
	// factorial := 1
	// n := 5
	// for i := 1; i <= n; i++ {
	// 	factorial = factorial * i
	// }
	// fmt.Println("this is factorial>>>>", factorial)

	// prime := 9

	// for i := 2; i <= prime; i++ {

	// 	if i%prime == 0 {
	// 		fmt.Println("this is  prime number", i)
	// 	}

	// }

	runfor := 10
	a := 0
	b := 1

	for i := 1; i <= runfor; i++ {
		fmt.Println(a, ">")

		c := a + b
		a = b
		b = c

	}

}
