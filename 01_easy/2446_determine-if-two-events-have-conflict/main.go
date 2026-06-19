package main

// LeetCode #2446: Determine if Two Events Have Conflict
// https://leetcode.com/problems/determine-if-two-events-have-conflict/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"01:15", "02:00"}, []string{"02:00", "03:00"})) // true
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"01:00", "02:00"}, []string{"01:20", "03:00"})) // true
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"10:00", "11:00"}, []string{"14:00", "15:00"})) // false
}

func timeToMin(t string) int {
	return int(t[0]-'0')*600 + int(t[1]-'0')*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}

func DetermineIfTwoEventsHaveConflict(event1 []string, event2 []string) bool {
	s1, e1 := timeToMin(event1[0]), timeToMin(event1[1])
	s2, e2 := timeToMin(event2[0]), timeToMin(event2[1])
	return !(e1 < s2 || e2 < s1)
}
