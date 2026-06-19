package main

// LeetCode #3168: Minimum Number of Chairs in a Waiting Room
// https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumChairs
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("EEEE"))   // 4
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("ELELEL")) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumChairs
func MinimumNumberOfChairsInAWaitingRoom(s string) int {
	current := 0
	maxChairs := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			current++
			if current > maxChairs {
				maxChairs = current
			}
		} else {
			current--
		}
	}
	return maxChairs
}
