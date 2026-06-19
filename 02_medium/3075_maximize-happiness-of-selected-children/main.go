package main

// LeetCode #3075: Maximize Happiness of Selected Children
// https://leetcode.com/problems/maximize-happiness-of-selected-children/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
}

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	ans := int64(0)
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val > 0 {
			ans += int64(val)
		}
	}
	return ans
}
