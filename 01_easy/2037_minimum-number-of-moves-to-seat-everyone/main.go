package main

// LeetCode #2037: Minimum Number of Moves to Seat Everyone
// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{3, 1, 5}, []int{2, 7, 4}))   // 4
	fmt.Println(MinimumNumberOfMovesToSeatEveryone([]int{4, 1, 5, 9}, []int{1, 3, 2, 6})) // 7
}

// Time: O(n log n), Space: O(1)
func MinimumNumberOfMovesToSeatEveryone(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	moves := 0
	for i := 0; i < len(seats); i++ {
		diff := seats[i] - students[i]
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
