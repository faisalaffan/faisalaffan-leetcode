package main

// LeetCode #525: Contiguous Array
// https://leetcode.com/problems/contiguous-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaxLength([]int{0, 1}))
	fmt.Println(FindMaxLength([]int{0, 1, 0}))
}

func FindMaxLength(nums []int) int {
	// Map count -> first index. count = (#1 - #0)
	countMap := make(map[int]int)
	countMap[0] = -1
	count := 0
	maxLen := 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if prevIdx, ok := countMap[count]; ok {
			if i-prevIdx > maxLen {
				maxLen = i - prevIdx
			}
		} else {
			countMap[count] = i
		}
	}

	return maxLen
}
