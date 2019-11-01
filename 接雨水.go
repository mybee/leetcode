package main

import "math"

func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	max := 0
	lmax := 0
	rmax := 0
	i := 0
	j := len(height) - 1
	for i < j {
		lmax = Max(lmax, height[i])
		rmax = Max(rmax, height[j])
		if lmax < rmax {
			max += lmax - height[i]
			i++
		} else {
			max += rmax - height[j]
			j--
		}
	}
	return max
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
