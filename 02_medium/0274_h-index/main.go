package main

// LeetCode #274: H-Index
// https://leetcode.com/problems/h-index/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{0}))
}
