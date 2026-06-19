package main

// LeetCode #1851: Minimum Interval to Include Each Query
// https://leetcode.com/problems/minimum-interval-to-include-each-query/
// Difficulty: Hard
//
// Given intervals [left, right] and queries, for each query find the minimum
// interval length (right-left+1) that contains it, or -1 if none.
//
// Approach: sort intervals by length ascending, sort queries with index.
// Use DSU (union-find) to efficiently skip already-answered queries.
// For each interval in increasing length, find all queries within [left, right]
// that haven't been answered yet and set their answer.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1:
	// intervals = [[1,4],[2,4],[3,6],[4,4]]
	// queries = [2,3,4,5]
	// Expected: [3,3,1,4]
	fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))

	// Example 2:
	// intervals = [[2,3],[2,5],[1,8],[20,25]]
	// queries = [2,19,5,22]
	// Expected: [2,-1,4,6]
	fmt.Println(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}))

	// Single interval, single query
	fmt.Println(minInterval([][]int{{1, 5}}, []int{3}))

	// No interval contains query
	fmt.Println(minInterval([][]int{{1, 2}, {5, 6}}, []int{4}))

	// Multiple queries, overlapping intervals
	fmt.Println(minInterval([][]int{{1, 10}, {2, 8}, {3, 6}}, []int{1, 4, 7, 10}))

	// Single query
	fmt.Println(minInterval([][]int{{1, 3}, {2, 5}, {1, 100}}, []int{4}))
}

func minInterval(intervals [][]int, queries []int) []int {
	n := len(queries)

	// Sort queries with their original indices
	type qi struct {
		val int
		idx int
	}
	qSorted := make([]qi, n)
	for i, q := range queries {
		qSorted[i] = qi{val: q, idx: i}
	}
	sort.Slice(qSorted, func(i, j int) bool {
		return qSorted[i].val < qSorted[j].val
	})

	// Sort intervals by length ascending
	sort.Slice(intervals, func(i, j int) bool {
		li := intervals[i][1] - intervals[i][0] + 1
		lj := intervals[j][1] - intervals[j][0] + 1
		return li < lj
	})

	// DSU: parent[i] = next unprocessed query index (or i itself)
	parent := make([]int, n+1)
	for i := 0; i <= n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	for _, iv := range intervals {
		l, r := iv[0], iv[1]
		length := r - l + 1

		// Find first query >= l
		start := sort.Search(n, func(i int) bool {
			return qSorted[i].val >= l
		})
		if start == n {
			continue
		}

		// Process all unassigned queries within [l, r]
		for idx := find(start); idx < n && qSorted[idx].val <= r; idx = find(idx) {
			ans[qSorted[idx].idx] = length
			parent[idx] = idx + 1
		}
	}
	return ans
}
