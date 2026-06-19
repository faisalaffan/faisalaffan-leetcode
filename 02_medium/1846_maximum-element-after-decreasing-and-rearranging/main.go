package main

// LeetCode #1846: Maximum Element After Decreasing and Rearranging
// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{100, 1, 1000}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{1, 2, 3, 4, 5}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func MaximumElementAfterDecreasingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}
