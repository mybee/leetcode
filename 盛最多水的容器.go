package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func maxArea(height []int) int {
	maxarea, l, r := 0, 0, len(height)-1
	for  {
		if l < r {
			min := float64(math.Min(float64(height[l]), float64(height[r] * (r - l))))
			fmt.Println(min)
			maxarea = int(math.Max(float64(maxarea), min))
			fmt.Println(maxarea)
			if height[l] < height[r] {
				l++
			} else {
				r--
			}
		} else {
			goto RETURN
		}
	}
RETURN:
	return maxarea
}