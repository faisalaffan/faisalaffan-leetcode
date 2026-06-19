package main

// LeetCode #3516: Find Closest Person
// https://leetcode.com/problems/find-closest-person/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestPerson(1, 2, 3))
	fmt.Println(FindClosestPerson(1, 3, 2))
	fmt.Println(FindClosestPerson(2, 1, 3))
}

// FindClosestPerson finds who is closer to person z. Returns 0 for tie, 1 for person 1, 2 for person 2.
// Time: O(1). Space: O(1).
func FindClosestPerson(x int, y int, z int) int {
	dx := z - x
	if dx < 0 {
		dx = -dx
	}
	dy := z - y
	if dy < 0 {
		dy = -dy
	}
	if dx < dy {
		return 1
	} else if dy < dx {
		return 2
	}
	return 0
}
