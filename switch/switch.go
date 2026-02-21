package main

import (
	"fmt"
	"time"
)

func main() {
	// simple
	i := 5
	switch i {
	case 1:
		fmt.Println("This is 1")
	case 5:
		fmt.Println("this is 2")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weeknd")

	default:
		fmt.Println("its work day")
	}

	// type switch
	whoAmI := func(i interface{}) {

		switch t := i.(type) {
		case int:
			fmt.Println("This is an integar")

		case string:
			fmt.Println("this is string>>>>>>")

		case bool:
			fmt.Println("Tis is bolean")

		default:
			fmt.Println("other", t)
		}

	}

	whoAmI("hello")
	whoAmI(2)
	whoAmI(true)
}
