package main

// LeetCode #2175: The Change in Global Rankings
// https://leetcode.com/problems/the-change-in-global-rankings/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func globalRankings(pointsBefore []int, pointsAfter []int) []int {
	n := len(pointsBefore)

	// Create sorted list of (points, originalIndex) for before
	type pair struct {
		points int
		idx    int
	}
	before := make([]pair, n)
	for i := 0; i < n; i++ {
		before[i] = pair{pointsBefore[i], i}
	}
	sort.Slice(before, func(i, j int) bool {
		if before[i].points != before[j].points {
			return before[i].points > before[j].points
		}
		return before[i].idx < before[j].idx
	})

	// Compute rank before: rank = position (1-indexed) when sorted descending
	rankBefore := make([]int, n)
	for pos, p := range before {
		rankBefore[p.idx] = pos + 1
	}

	// Create sorted list for after
	after := make([]pair, n)
	for i := 0; i < n; i++ {
		after[i] = pair{pointsAfter[i], i}
	}
	sort.Slice(after, func(i, j int) bool {
		if after[i].points != after[j].points {
			return after[i].points > after[j].points
		}
		return after[i].idx < after[j].idx
	})

	// Compute rank after
	rankAfter := make([]int, n)
	for pos, p := range after {
		rankAfter[p.idx] = pos + 1
	}

	// Difference
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rankBefore[i] - rankAfter[i]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", globalRankings([]int{10, 20, 30}, []int{15, 25, 35}))
	// Expected: [0, 0, 0] (same relative order)

	// Test case 2
	fmt.Println("Test 2:", globalRankings([]int{50, 40, 30, 20}, []int{45, 45, 35, 25}))
	// Expected: varying based on rank changes
}
