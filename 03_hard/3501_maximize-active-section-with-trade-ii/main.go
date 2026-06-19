package main

// LeetCode #3501: Maximize Active Section with Trade II
// https://leetcode.com/problems/maximize-active-section-with-trade-ii/
// Difficulty: Hard
//
// Given a binary string s, for each query [l, r], consider the substring
// s[l..r] augmented with '1' at both ends. You can perform one trade:
// select a '0' and turn it into '1' (activating a section). Maximize the
// number of active sections (contiguous '1's) after the trade.
//
// Approach: Precompute segment runs and for each query, find the maximum
// possible active sections after merging adjacent segments.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxActiveSectionsAfterTrade("1001", [][]int{{0, 3}}))
	// Example 2
	fmt.Println(maxActiveSectionsAfterTrade("10101", [][]int{{0, 4}, {1, 3}}))
	// Example 3: all zeros
	fmt.Println(maxActiveSectionsAfterTrade("000", [][]int{{0, 2}}))
	// Edge: single char
	fmt.Println(maxActiveSectionsAfterTrade("1", [][]int{{0, 0}}))
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)

	// Precompute prefix sums of '1's
	prefOne := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefOne[i+1] = prefOne[i]
		if s[i] == '1' {
			prefOne[i+1]++
		}
	}

	// Find segments of '0's
	type segment struct {
		l, r int // inclusive
	}
	var zeros []segment
	i := 0
	for i < n {
		if s[i] == '0' {
			start := i
			for i < n && s[i] == '0' {
				i++
			}
			zeros = append(zeros, segment{start, i - 1})
		} else {
			i++
		}
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]

		// Count '1's in the query range
		ones := prefOne[r+1] - prefOne[l]
		if ones == 0 {
			// All zeros -> can make 1 active section by flipping one 0
			ans[qi] = 1
			continue
		}

		// Base active sections = number of '1' runs in query range
		// Flipping a '0' can merge adjacent '1' sections
		// Find which zero segment, when flipped, produces the most new sections

		base := ones // at least each '1' is its own section (worst case)
		best := base

		for _, seg := range zeros {
			if seg.l > r || seg.r < l {
				continue
			}
			// This zero segment overlaps with query range
			zeroLen := min(seg.r, r) - max(seg.l, l) + 1
			if zeroLen <= 0 {
				continue
			}

			// Count '1's adjacent to this zero segment
			adjLeft := 0
			if seg.l-1 >= l && s[seg.l-1] == '1' {
				leftRunEnd := seg.l - 1
				for leftRunEnd >= l && s[leftRunEnd] == '1' {
					adjLeft++
					leftRunEnd--
				}
			}
			adjRight := 0
			if seg.r+1 <= r && s[seg.r+1] == '1' {
				rightRunStart := seg.r + 1
				for rightRunStart <= r && s[rightRunStart] == '1' {
					adjRight++
					rightRunStart++
				}
			}

			// Flipping this zero merges adjacent 1 sections
			// Gains: the zero becomes 1, and if it merges sections, we reduce section count
			if adjLeft > 0 && adjRight > 0 {
				// Merges two 1 sections -> we reduce section count by 1
				// But we add 1 for the flipped zero
				result := ones + adjLeft + adjRight - 1 // merged: was 2 sections, now 1
				if result > best {
					best = result
				}
			} else if adjLeft > 0 || adjRight > 0 {
				result := ones + adjLeft + adjRight
				if result > best {
					best = result
				}
			}
		}

		ans[qi] = best
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
