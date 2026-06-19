package main

// LeetCode #3346: Maximum Frequency of an Element After Performing Operations I
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-i/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2)) // 2
	fmt.Println(maxFrequency([]int{5, 11, 20}, 5, 3)) // 2
}

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)

	// Count duplicates (existing frequency)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Sliding window: for each value in sorted array,
	// find how many numbers can reach it within k operations
	n := len(nums)
	ans := 0
	left := 0
	right := 0

	// Collect unique values
	unique := make([]int, 0, len(freq))
	for v := range freq {
		unique = append(unique, v)
	}
	sort.Ints(unique)

	for _, v := range unique {
		// Expand right: all nums[right] <= v + k
		for right < n && nums[right] <= v+k {
			right++
		}
		// Expand left: all nums[left] < v - k
		for left < n && nums[left] < v-k {
			left++
		}
		total := right - left
		candidates := total - freq[v]
		ops := numOperations
		if candidates > ops {
			candidates = ops
		}
		cur := freq[v] + candidates
		if cur > ans {
			ans = cur
		}
	}

	// Also consider values not in nums (midpoints)
	// For any pair of values where the distance is <= 2*k,
	// operations can be split between them
	left = 0
	for right = 0; right < n; right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		total := right - left + 1
		if total > numOperations {
			total = numOperations
		}
		if total > ans {
			ans = total
		}
	}

	return ans
}
