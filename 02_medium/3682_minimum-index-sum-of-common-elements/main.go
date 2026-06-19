package main

// LeetCode #3682: Minimum Index Sum of Common Elements
// https://leetcode.com/problems/minimum-index-sum-of-common-elements/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func minimumIndexSumOfCommonElements(nums1 []int, nums2 []int) int {
	idxMap := make(map[int]int)
	for i, v := range nums2 {
		idxMap[v] = i
	}

	minSum := int(^uint(0) >> 1)
	found := false

	for i, v := range nums1 {
		if j, ok := idxMap[v]; ok {
			sum := i + j
			if !found || sum < minSum {
				minSum = sum
				found = true
			}
		}
	}

	if !found {
		return -1
	}
	return minSum
}

func main() {
	fmt.Println(minimumIndexSumOfCommonElements([]int{3, 2, 1}, []int{1, 3, 1}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{5, 1, 2}, []int{2, 1, 3}))
	fmt.Println(minimumIndexSumOfCommonElements([]int{6, 4}, []int{7, 8}))
}
