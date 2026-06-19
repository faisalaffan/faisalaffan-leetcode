package main

// LeetCode #3140: Consecutive Available Seats II
// https://leetcode.com/problems/consecutive-available-seats-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func consecutiveAvailableSeats(seats [][]int) int {
	sort.Slice(seats, func(i, j int) bool {
		if seats[i][0] != seats[j][0] {
			return seats[i][0] < seats[j][0]
		}
		return seats[i][1] < seats[j][1]
	})

	maxLen := 0
	curStart, curEnd := seats[0][0], seats[0][1]

	for i := 1; i < len(seats); i++ {
		if seats[i][0] <= curEnd+1 {
			if seats[i][1] > curEnd {
				curEnd = seats[i][1]
			}
		} else {
			if curEnd-curStart+1 > maxLen {
				maxLen = curEnd - curStart + 1
			}
			curStart, curEnd = seats[i][0], seats[i][1]
		}
	}
	if curEnd-curStart+1 > maxLen {
		maxLen = curEnd - curStart + 1
	}

	return maxLen
}

func main() {
	fmt.Println(consecutiveAvailableSeats([][]int{{1, 3}, {4, 6}, {7, 10}})) // Expected: 10
	fmt.Println(consecutiveAvailableSeats([][]int{{1, 2}, {5, 6}}))          // Expected: 2
}
