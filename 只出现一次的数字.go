package main

import "fmt"

func main() {
	a := []int{2, 4, 4, 4}
	b := singleNumber(a)
	fmt.Println(b)
}

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}