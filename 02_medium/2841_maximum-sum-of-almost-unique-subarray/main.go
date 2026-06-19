package main

// LeetCode #2841: Maximum Sum of Almost Unique Subarray
// https://leetcode.com/problems/maximum-sum-of-almost-unique-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumSumOfAlmostUniqueSubarray(nums []int, m int, k int) int64 {
	n := len(nums)
	if n < k {
		return 0
	}

	freq := make(map[int]int)
	var sum int64
	var best int64

	for i := 0; i < k; i++ {
		freq[nums[i]]++
		sum += int64(nums[i])
	}

	if len(freq) >= m {
		best = sum
	}

	for i := k; i < n; i++ {
		// Remove leftmost
		left := nums[i-k]
		freq[left]--
		if freq[left] == 0 {
			delete(freq, left)
		}
		sum -= int64(left)

		// Add rightmost
		right := nums[i]
		freq[right]++
		sum += int64(right)

		if len(freq) >= m && sum > best {
			best = sum
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{2, 6, 7, 3, 1, 7}, 3, 4))
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{1, 1, 1, 3}, 2, 2))
}
