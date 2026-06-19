package main

// LeetCode #3649: Number of Perfect Pairs
// https://leetcode.com/problems/number-of-perfect-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func numberOfPerfectPairs(nums []int) int64 {
	n := len(nums)
	arr := make([]int, n)
	for i, v := range nums {
		if v < 0 {
			arr[i] = -v
		} else {
			arr[i] = v
		}
	}

	sort.Ints(arr)

	var ans int64 = 0
	j := 0
	for i := 1; i < n; i++ {
		for 2*arr[j] < arr[i] {
			j++
		}
		ans += int64(i - j)
	}
	return ans
}

func main() {
	fmt.Println(numberOfPerfectPairs([]int{1, 2, 3, 4}))
	fmt.Println(numberOfPerfectPairs([]int{-1, 1, -2, 2}))
	fmt.Println(numberOfPerfectPairs([]int{5, 1, 2}))
}
