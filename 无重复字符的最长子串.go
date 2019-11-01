package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	lastOccuerd := make(map[rune]int)
	start, length := 0, 0
	for i, v := range s {
		if last, ok := lastOccuerd[v]; ok && last >= start {
			start = last + 1
		}
		//判断当前是否最长
		if i-start+1 > length {
			length = i - start + 1
		}
		// 记录lastOc
		lastOccuerd[v] = i
	}
	return length
}
