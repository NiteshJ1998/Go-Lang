package main

import "fmt"

func main(){
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2)
	// capacity -> maximum numbers of element can fit
	fmt.Println(cap(nums))
	fmt.Println(nums)
}
