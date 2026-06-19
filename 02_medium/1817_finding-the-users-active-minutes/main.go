package main

// LeetCode #1817: Finding the Users Active Minutes
// https://leetcode.com/problems/finding-the-users-active-minutes/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userMinutes := make(map[int]map[int]bool)
	for _, log := range logs {
		id, min := log[0], log[1]
		if userMinutes[id] == nil {
			userMinutes[id] = make(map[int]bool)
		}
		userMinutes[id][min] = true
	}

	result := make([]int, k)
	for _, minutes := range userMinutes {
		uam := len(minutes)
		if uam <= k {
			result[uam-1]++
		}
	}
	return result
}

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5)) // Expected: [0, 2, 0, 0, 0]
	fmt.Println(findingUsersActiveMinutes([][]int{{1, 1}, {2, 2}, {2, 3}}, 4)) // Expected: [1, 1, 0, 0]
}
