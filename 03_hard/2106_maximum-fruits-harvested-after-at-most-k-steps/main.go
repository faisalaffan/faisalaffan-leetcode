package main

// LeetCode #2106: Maximum Fruits Harvested After at Most K Steps
// https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/
// Difficulty: Hard
//
// Approach: Prefix sum + sliding window.
// Build a fruit amount array up to max position (200000).
// Use prefix sums to query range sums in O(1).
// For each possible left steps (0..k), compute remaining steps for right,
// and vice versa. Take max of all ranges.

import "fmt"

func main() {
	// Example from problem statement
	fruits1 := [][]int{{2, 8}, {6, 3}, {8, 6}}
	startPos1 := 5
	k1 := 4
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d (expected 9)\n",
		fruits1, startPos1, k1, maxTotalFruits(fruits1, startPos1, k1))

	// Additional tests
	fruits2 := [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}
	startPos2 := 5
	k2 := 4
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d (expected 14)\n",
		fruits2, startPos2, k2, maxTotalFruits(fruits2, startPos2, k2))

	fruits3 := [][]int{{0, 3}, {6, 4}, {8, 5}}
	startPos3 := 3
	k3 := 2
	fmt.Printf("maxTotalFruits(%v, %d, %d) = %d\n",
		fruits3, startPos3, k3, maxTotalFruits(fruits3, startPos3, k3))
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	const maxPos = 200000
	amount := make([]int, maxPos+1)
	for _, f := range fruits {
		amount[f[0]] = f[1]
	}

	prefix := make([]int, maxPos+2)
	for i := 0; i <= maxPos; i++ {
		prefix[i+1] = prefix[i] + amount[i]
	}

	sumRange := func(l, r int) int {
		if l < 0 {
			l = 0
		}
		if r > maxPos {
			r = maxPos
		}
		if l > r {
			return 0
		}
		return prefix[r+1] - prefix[l]
	}

	ans := 0
	// Go left first, then right
	for left := 0; left <= k; left++ {
		right := max(0, k-2*left)
		l := startPos - left
		r := startPos + right
		ans = max(ans, sumRange(l, r))
	}

	// Go right first, then left
	for right := 0; right <= k; right++ {
		left := max(0, k-2*right)
		l := startPos - left
		r := startPos + right
		ans = max(ans, sumRange(l, r))
	}

	return ans
}
