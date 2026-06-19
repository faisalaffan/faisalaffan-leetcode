package main

// LeetCode #2948: Make Lexicographically Smallest Array by Swapping Elements
// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 5, 3, 9, 8}, 2))
	fmt.Println(lexicographicallySmallestArray([]int{1, 7, 6, 18, 2, 1}, 3))
	fmt.Println(lexicographicallySmallestArray([]int{1, 2, 3}, 0))
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})
	ans := make([]int, n)
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[idx[j]]-nums[idx[j-1]] <= limit {
			j++
		}
		t := make([]int, j-i)
		copy(t, idx[i:j])
		sort.Ints(t)
		for k := i; k < j; k++ {
			ans[t[k-i]] = nums[idx[k]]
		}
		i = j
	}
	return ans
}
