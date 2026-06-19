package main

// LeetCode #2945: Find Maximum Non-decreasing Array Length
// https://leetcode.com/problems/find-maximum-non-decreasing-array-length/
// Difficulty: Hard
//
// Approach: DP + monotonic deque optimization.
// We partition the array into contiguous groups, replace each group
// with its sum, and want the resulting array to be non-decreasing.
// Goal: maximize the number of groups.
//
// Define:
//   f[i] = max groups for prefix ending at i-1 (i elements total)
//   g[i] = minimum possible last group sum achieving f[i]
//   pref[i] = prefix sum of first i elements
//
// Transition: for j < i where sum(j..i-1) = pref[i]-pref[j] >= g[j],
//   f[i] = f[j] + 1, g[i] = pref[i]-pref[j]
//
// Optimization: maintain deque of candidates sorted by g[j] and pref[j]+g[j].

import (
	"fmt"
)

func findMaximumLength(nums []int) int {
	n := len(nums)
	pref := make([]int64, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + int64(v)
	}

	f := make([]int, n+1)
	g := make([]int64, n+1) // g[j] = last group sum in optimal partition of first j elements
	deq := make([]int, 0, n+1)
	deq = append(deq, 0) // start with index 0

	head := 0
	for i := 1; i <= n; i++ {
		// Pop front: discard indices that are no longer optimal
		// A candidate j is outdated if pref[i] >= pref[deq[head+1]] + g[deq[head+1]]
		for head+1 < len(deq) && pref[i] >= pref[deq[head+1]]+g[deq[head+1]] {
			head++
		}

		j := deq[head]
		f[i] = f[j] + 1
		last := pref[i] - pref[j]
		g[i] = last

		// Pop back: maintain monotonicity of pref[i] + g[i] (strictly increasing)
		for len(deq) > head && pref[i]+g[i] <= pref[deq[len(deq)-1]]+g[deq[len(deq)-1]] {
			deq = deq[:len(deq)-1]
		}
		deq = append(deq, i)
	}
	return f[n]
}

func main() {
	// Example: [2,3,1,4,5] -> 4
	// Partition: [2],[3],[1,4],[5] -> sums [2,3,5,5] (non-decreasing, length 4)
	fmt.Println(findMaximumLength([]int{2, 3, 1, 4, 5}))

	// Simple cases
	fmt.Println(findMaximumLength([]int{1, 2, 3}))
	fmt.Println(findMaximumLength([]int{5, 4, 3, 2, 1}))
}
