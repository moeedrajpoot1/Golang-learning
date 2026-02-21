package main

import (
	"fmt"
)

func main() {

	// we can say dynamically arrays
	// most used constrcut in go
	// useful methods

	// un intialized slice is nill

	// var nums []int

	// fmt.Println("nums>>>>", nums == nil)

	// var nums = make([]int, 2, 5)

	// fmt.Println("numssss", nums == nil)

	// fmt.Println(cap(nums))

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)

	// fmt.Println("numsss", nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// var nums = []int{1, 2, 3, 4, 5, 6}

	// fmt.Println("slice", nums[0:2])
	// fmt.Println("slice", nums[:1])
	// fmt.Println("slice", nums[1:])
	// fmt.Println("slice", nums[1:])

	// slice package

	// var nums = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 3, 4}

	// fmt.Println(slices.Equal(nums, nums2))

	// 2D slices

	var nums = [][]int{{1, 2}, {3, 4, 5}}

	fmt.Println("this is 2 D Slices", nums)

}
