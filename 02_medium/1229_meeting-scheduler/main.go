package main

import (
	"fmt"
	"sort"
)

// LeetCode #1229: Meeting Scheduler
// https://leetcode.com/problems/meeting-scheduler/
// Difficulty: Medium [Paid]

// Find earliest time slot of given duration that works for both.
// Two pointers approach after sorting slots by start time.

// Time: O(n log n + m log m)
// Space: O(1)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
	sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })

	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		start := max(slots1[i][0], slots2[j][0])
		end := min(slots1[i][1], slots2[j][1])
		if end-start >= duration {
			return []int{start, start + duration}
		}
		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("%v (expected: [60 68])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}},
			[][]int{{0, 15}, {60, 70}}, 8))

	fmt.Printf("%v (expected: [])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}},
			[][]int{{0, 15}, {55, 58}}, 5))
}
