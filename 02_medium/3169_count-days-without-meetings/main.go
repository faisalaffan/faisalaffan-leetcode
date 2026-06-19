package main

// LeetCode #3169: Count Days Without Meetings
// https://leetcode.com/problems/count-days-without-meetings/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	merged := make([][2]int, 0)
	for _, m := range meetings {
		if len(merged) > 0 && m[0] <= merged[len(merged)-1][1]+1 {
			if m[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = m[1]
			}
		} else {
			merged = append(merged, [2]int{m[0], m[1]})
		}
	}

	ans := days
	for _, m := range merged {
		ans -= m[1] - m[0] + 1
	}
	return ans
}

func main() {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})) // Expected: 2
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))            // Expected: 1
	fmt.Println(countDays(6, [][]int{{1, 6}}))                    // Expected: 0
}
