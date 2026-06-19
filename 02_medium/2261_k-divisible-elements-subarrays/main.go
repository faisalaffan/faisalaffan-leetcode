package main

// LeetCode #2261: K Divisible Elements Subarrays
// https://leetcode.com/problems/k-divisible-elements-subarrays/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n^2)

import (
	"fmt"
	"strconv"
	"strings"
)

func countDistinct(nums []int, k int, p int) int {
	n := len(nums)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
		count := 0
		var sb strings.Builder
		for j := i; j < n; j++ {
			if nums[j]%p == 0 {
				count++
			}
			if count > k {
				break
			}
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(nums[j]))
			seen[sb.String()] = true
		}
	}
	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println(countDistinct([]int{2, 3, 3, 2, 2}, 2, 2))
	// Expected: 11

	// Test case 2
	fmt.Println(countDistinct([]int{1, 2, 3, 4}, 4, 5))
	// Expected: 10
}
