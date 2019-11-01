package main

import "fmt"

func main() {
	a := twoSumOneHash([]int{3,2,4}, 6)
	fmt.Println(a)
}
// 暴力算法
func twoSum(nums []int, target int) []int {
	for i:=0; i< len(nums); i++{
		for j:= i+1; j < len(nums); j++ {
			fmt.Println(i, j)
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 两遍hash法
func twoSumTwoHash(nums []int, target int) []int {
	return []int{}
}

// 一遍hash法
func twoSumOneHash(nums []int, target int) []int {
	intMap := make(map[int] int)
	for i:= 0; i<len(nums); i++ {
		if j, ok := intMap[target - nums[i]]; ok {
			return []int{i, j}
		}
		intMap[nums[i]] = i
		fmt.Println(intMap)
	}
	return []int{}
}
