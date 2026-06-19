package main

// LeetCode #1982: Find Array Given Subset Sums
// https://leetcode.com/problems/find-array-given-subset-sums/
// Difficulty: Hard
// Approach: Sort + recursive extraction.
// For sorted sums, the difference between consecutive elements gives a candidate.
// Partition into "without" and "with" the candidate. Recurse on "without".

import (
	"fmt"
	"sort"
)

func recoverArray(n int, sums []int) []int {
	sort.Ints(sums)
	ans := make([]int, 0, n)

	for len(sums) > 1 {
		// Candidate: diff between smallest two elements
		diff := sums[1] - sums[0]

		// Partition sums into left (without diff) and right (with diff)
		freq := make(map[int]int)
		for _, v := range sums {
			freq[v]++
		}

		left := make([]int, 0, len(sums)/2)
		right := make([]int, 0, len(sums)/2)
		for _, v := range sums {
			if freq[v] == 0 {
				continue
			}
			freq[v]--
			freq[v+diff]--
			left = append(left, v)
			right = append(right, v+diff)
		}

		// Check if left contains 0 (the empty set sum)
		// If yes, diff is positive (or 0 which shouldn't happen for valid input)
		// If no, diff is negative (right contains 0 instead)
		hasZero := false
		for _, v := range left {
			if v == 0 {
				hasZero = true
				break
			}
		}

		if !hasZero {
			left, right = right, left
			diff = -diff
		}

		ans = append(ans, diff)
		sums = left
	}

	return ans
}

func main() {
	// Example: n=3, sums=[-3,-2,-1,0,0,1,2,3] -> [1,2,-3] or [1,-2,3] etc.
	// The problem states sums are distinct, but standard LeetCode example:
	// Let's use a proper test case
	fmt.Println(recoverArray(3, []int{-3, -2, -1, 0, 0, 1, 2, 3}))

	// n=2, sums=[0,1,2,3] -> original array sum = 3, subset sums = {0, a, b, a+b}
	// If a=1, b=2: sums={0,1,2,3}. So result is [1,2] or [2,1]
	fmt.Println(recoverArray(2, []int{0, 1, 2, 3}))

	// Another test
	fmt.Println(recoverArray(2, []int{0, 1, 1, 2}))
}
