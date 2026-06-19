package main

// LeetCode #354: Russian Doll Envelopes
// https://leetcode.com/problems/russian-doll-envelopes/
// Difficulty: Hard
//
// Sort envelopes by width ascending. When widths are equal, sort by height
// descending to prevent same-width nesting. Then find LIS (Longest Increasing
// Subsequence) on heights using patience sorting O(n log n).

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: [[5,4],[6,4],[6,7],[2,3]] -> 3 (2,3 -> 5,4 -> 6,7)
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	// Example 2: [[1,1],[1,1],[1,1]] -> 1
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
	// Example 3: [[1,2],[2,3],[3,4],[4,5]] -> 4
	fmt.Println(maxEnvelopes([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	// Edge: single envelope
	fmt.Println(maxEnvelopes([][]int{{1, 1}}))
	// Edge: all same width
	fmt.Println(maxEnvelopes([][]int{{1, 2}, {1, 3}, {1, 4}}))
}

func maxEnvelopes(envelopes [][]int) int {
	// Sort by width ascending; if width ties, height descending
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	// LIS on heights using patience sorting (binary search)
	tails := make([]int, 0, len(envelopes))
	for _, e := range envelopes {
		h := e[1]
		idx := sort.SearchInts(tails, h)
		if idx == len(tails) {
			tails = append(tails, h)
		} else {
			tails[idx] = h
		}
	}
	return len(tails)
}
