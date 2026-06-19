package main

// LeetCode #252: Meeting Rooms
// https://leetcode.com/problems/meeting-rooms/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(CanAttendMeetings([][]int{{7, 10}, {2, 4}}))
}
