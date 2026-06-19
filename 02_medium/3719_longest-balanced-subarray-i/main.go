package main

// LeetCode #3719: Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func longestBalancedSubarrayI(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		if n-i <= ans {
			break
		}
		evenVisited := make(map[int]bool)
		oddVisited := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenVisited[nums[j]] {
					evenVisited[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddVisited[nums[j]] {
					oddVisited[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubarrayI([]int{1, 2, 3, 4}))
	fmt.Println(longestBalancedSubarrayI([]int{2, 4, 6, 8}))
	fmt.Println(longestBalancedSubarrayI([]int{1, 1, 2, 2, 3, 3}))
}
