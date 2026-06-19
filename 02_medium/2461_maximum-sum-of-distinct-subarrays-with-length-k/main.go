package main

// LeetCode #2461: Maximum Sum of Distinct Subarrays With Length K
// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Sliding window with frequency map for distinct check.

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)) // 15
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))              // 0
}

func maximumSubarraySum(nums []int, k int) int64 {
	freq := make(map[int]int)
	var sum, ans int64
	dupCount := 0

	for i, v := range nums {
		sum += int64(v)
		freq[v]++
		if freq[v] == 2 {
			dupCount++
		}

		if i >= k {
			left := nums[i-k]
			sum -= int64(left)
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
		}

		if i >= k-1 && dupCount == 0 && sum > ans {
			ans = sum
		}
	}
	return ans
}
