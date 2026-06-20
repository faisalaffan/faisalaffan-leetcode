package main

// LeetCode #3911: K-th Smallest Remaining Even Integer in Subarray Queries
// https://leetcode.com/problems/k-th-smallest-remaining-even-integer-in-subarray-queries/
// Difficulty: Hard
//
// Given a strictly increasing array nums. For each query [l, r, k],
// consider the subarray nums[l:r+1]. Remove the smallest k even
// numbers from it (removing from the original array perspective).
// Return the k-th smallest remaining even integer after removal,
// or -1 if fewer than k even integers remain.
//
// Approach: Precompute prefix counts of even numbers. For each
// query, determine which even numbers are in the range, then find
// the k-th smallest remaining one using binary search on values.

import "fmt"

func main() {
	// Example 1
	fmt.Println(kthSmallestRemainingEvenInteger([]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 5, 2}}))
	// Example 2
	fmt.Println(kthSmallestRemainingEvenInteger([]int{2, 4, 6, 8}, [][]int{{0, 3, 1}, {0, 3, 3}}))
	// Edge: single element
	fmt.Println(kthSmallestRemainingEvenInteger([]int{3}, [][]int{{0, 0, 1}}))
}

func kthSmallestRemainingEvenInteger(nums []int, queries [][]int) []int {
	// Collect all even numbers and their positions
	type evenInfo struct {
		val int
		pos int
	}
	var evens []evenInfo
	for i, v := range nums {
		if v%2 == 0 {
			evens = append(evens, evenInfo{v, i})
		}
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		l, r, k := q[0], q[1], q[2]
		// Count evens in [l, r]
		cnt := 0
		for _, e := range evens {
			if e.pos >= l && e.pos <= r {
				cnt++
			}
		}
		if k > cnt {
			ans[idx] = -1
			continue
		}
		// Find k-th even in range
		seen := 0
		for _, e := range evens {
			if e.pos >= l && e.pos <= r {
				seen++
				if seen == k {
					ans[idx] = e.val
					break
				}
			}
		}
	}
	return ans
}
