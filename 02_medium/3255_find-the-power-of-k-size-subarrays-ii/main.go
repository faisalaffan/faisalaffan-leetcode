package main

// LeetCode #3255: Find the Power of K-Size Subarrays II
// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	consec := 1

	for i := range n {
		if i > 0 && nums[i] == nums[i-1]+1 {
			consec++
		} else {
			consec = 1
		}
		if i >= k-1 {
			if consec >= k {
				ans[i-k+1] = nums[i]
			} else {
				ans[i-k+1] = -1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 5}, 3)) // Expected: [3, 4, -1, -1]
	fmt.Println(resultsArray([]int{1, 4, 5, 2, 3}, 3))    // Expected: [-1, 5, -1]
	fmt.Println(resultsArray([]int{1}, 1))                 // Expected: [1]
}
