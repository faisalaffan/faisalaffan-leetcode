package main

// LeetCode #523: Continuous Subarray Sum
// https://leetcode.com/problems/continuous-subarray-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(k) where k = min(k, n)

import "fmt"

func main() {
	fmt.Println(CheckSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func CheckSubarraySum(nums []int, k int) bool {
	// Map remainder -> first index
	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	sum := 0

	for i, num := range nums {
		sum += num
		rem := sum % k
		if rem < 0 {
			rem += k
		}
		if prevIdx, ok := remainderMap[rem]; ok {
			if i-prevIdx >= 2 {
				return true
			}
		} else {
			remainderMap[rem] = i
		}
	}

	return false
}
