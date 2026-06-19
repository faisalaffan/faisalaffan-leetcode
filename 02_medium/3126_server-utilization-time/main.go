package main

// LeetCode #3126: Server Utilization Time
// https://leetcode.com/problems/server-utilization-time/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func serverUtilizationTime(logs [][]int) int {
	if len(logs) == 0 {
		return 0
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	merged := [][]int{logs[0]}
	for i := 1; i < len(logs); i++ {
		last := merged[len(merged)-1]
		if logs[i][0] <= last[1] {
			if logs[i][1] > last[1] {
				last[1] = logs[i][1]
			}
		} else {
			merged = append(merged, logs[i])
		}
	}

	total := 0
	for _, seg := range merged {
		total += seg[1] - seg[0]
	}
	return total
}

func main() {
	fmt.Println(serverUtilizationTime([][]int{{0, 5}, {2, 7}, {8, 10}})) // Expected: 9
	fmt.Println(serverUtilizationTime([][]int{{1, 3}, {3, 5}}))          // Expected: 4
}
