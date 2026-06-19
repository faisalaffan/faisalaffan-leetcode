package main

// LeetCode #3410: Maximize Subarray Sum After Removing All Occurrences of One Element
// https://leetcode.com/problems/maximize-subarray-sum-after-removing-all-occurrences-of-one-element/
// Difficulty: Hard
//
// Kadane variant: skip all occurrences of one value.
// For each distinct value, run Kadane treating that value as 0.
// Also compute standard Kadane (no removal).

import "fmt"

func main() {
	fmt.Println(MaximizeSubarraySumAfterRemovingAllOccurrencesOfOneElement([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func MaximizeSubarraySumAfterRemovingAllOccurrencesOfOneElement(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Standard Kadane (no removal)
	best := int64(nums[0])
	cur := int64(0)
	for _, v := range nums {
		cur += int64(v)
		if cur > best {
			best = cur
		}
		if cur < 0 {
			cur = 0
		}
	}

	// Try removing each distinct value
	vals := make(map[int]bool)
	for _, v := range nums {
		vals[v] = true
	}

	for skipVal := range vals {
		cur = 0
		for _, v := range nums {
			if v == skipVal {
				continue
			}
			cur += int64(v)
			if cur > best {
				best = cur
			}
			if cur < 0 {
				cur = 0
			}
		}
	}
	return best
}
