package main

// LeetCode #2780: Minimum Index of a Valid Split
// https://leetcode.com/problems/minimum-index-of-a-valid-split/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumIndexOfAValidSplit(nums []int) int {
	n := len(nums)

	// Find majority element using Boyer-Moore
	majority := nums[0]
	count := 0
	for _, num := range nums {
		if count == 0 {
			majority = num
			count = 1
		} else if num == majority {
			count++
		} else {
			count--
		}
	}

	// Count total occurrences of majority
	totalCount := 0
	for _, num := range nums {
		if num == majority {
			totalCount++
		}
	}

	// Find minimum split index
	leftCount := 0
	for i := 0; i < n-1; i++ {
		if nums[i] == majority {
			leftCount++
		}
		rightCount := totalCount - leftCount
		leftLen := i + 1
		rightLen := n - i - 1
		if leftCount*2 > leftLen && rightCount*2 > rightLen {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(MinimumIndexOfAValidSplit([]int{1, 2, 2, 2}))
	fmt.Println(MinimumIndexOfAValidSplit([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
}
