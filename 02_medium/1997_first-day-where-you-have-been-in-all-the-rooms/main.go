package main

// LeetCode #1997: First Day Where You Have Been in All the Rooms
// https://leetcode.com/problems/first-day-where-you-have-been-in-all-the-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 1, 2, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0, 2}))
}

// Time: O(n), Space: O(n)
func FirstDayWhereYouHaveBeenInAllTheRooms(nextVisit []int) int {
	const mod = 1_000_000_007
	n := len(nextVisit)
	s := make([]int, n)

	for i := 0; i < n-1; i++ {
		j := nextVisit[i]
		s[i+1] = (s[i]*2 - s[j] + 2) % mod
		if s[i+1] < 0 {
			s[i+1] += mod
		}
	}

	return s[n-1]
}
