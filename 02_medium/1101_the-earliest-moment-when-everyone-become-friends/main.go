package main

// LeetCode #1101: The Earliest Moment When Everyone Become Friends
// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/
// Difficulty: Medium
//
// Approach: Sort logs by timestamp, Union-Find
// Time: O(n log n + m * alpha(n)) where n = logs, m = N
// Space: O(N)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestAcq([][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6)) // 20190301
	fmt.Println(earliestAcq([][]int{{0, 0, 1}, {1, 1, 2}, {2, 0, 2}}, 3)) // 2
}

func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	groups := n
	for _, log := range logs {
		ts, a, b := log[0], log[1], log[2]
		if find(a) != find(b) {
			union(a, b)
			groups--
			if groups == 1 {
				return ts
			}
		}
	}

	return -1
}
