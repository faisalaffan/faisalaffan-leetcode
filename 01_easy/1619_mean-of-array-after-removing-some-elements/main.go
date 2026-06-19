package main

// LeetCode #1619: Mean of Array After Removing Some Elements
// https://leetcode.com/problems/mean-of-array-after-removing-some-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func TrimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	remove := n / 20
	sum := 0
	for i := remove; i < n-remove; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n-2*remove)
}

func main() {
	fmt.Println(TrimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
	fmt.Println(TrimMean([]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}))
	fmt.Println(TrimMean([]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}))
}
