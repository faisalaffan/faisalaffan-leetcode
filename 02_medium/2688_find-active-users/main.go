package main

// LeetCode #2688: Find Active Users
// https://leetcode.com/problems/find-active-users/
// Difficulty: Medium [Paid] (SQL problem)
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findActiveUsers(logins [][]int) []int {
	// Group by user
	userLogins := make(map[int][]int)
	for _, log := range logins {
		user, day := log[0], log[1]
		userLogins[user] = append(userLogins[user], day)
	}

	active := []int{}
	for user, days := range userLogins {
		if len(days) < 5 {
			continue
		}
		sort.Ints(days)
		// Check for 5 consecutive days (days are consecutive if each diff is 1)
		for i := 0; i <= len(days)-5; i++ {
			if days[i+4]-days[i] == 4 {
				active = append(active, user)
				break
			}
		}
	}

	sort.Ints(active)
	return active
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findActiveUsers([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 1}, {2, 3}}))
	// Expected: [1]

	// Test case 2
	fmt.Println("Test 2:", findActiveUsers([][]int{{1, 1}, {1, 3}, {1, 5}}))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", findActiveUsers([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 7}, {2, 8}, {2, 9}, {2, 10}, {2, 11}}))
	// Expected: [1 2]
}
