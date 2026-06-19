package main

// LeetCode #1940: Longest Common Subsequence Between Sorted Arrays
// https://leetcode.com/problems/longest-common-subsequence-between-sorted-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{1, 3, 4}, {1, 4, 7, 9}}))
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}))
}

// Time: O(total elements), Space: O(unique elements)
func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int {
	// Since arrays are sorted, use frequency counting
	// Numbers appearing in ALL arrays are the answer
	freq := make(map[int]int)
	for _, arr := range arrs {
		for _, v := range arr {
			freq[v]++
		}
	}

	n := len(arrs)
	result := make([]int, 0)
	for _, v := range arrs[0] {
		if freq[v] == n {
			result = append(result, v)
		}
	}
	return result
}
