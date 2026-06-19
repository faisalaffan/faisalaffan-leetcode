package main

// LeetCode #1939: Users That Actively Request Confirmation Messages
// https://leetcode.com/problems/users-that-actively-request-confirmation-messages/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// user actions: each pair is (userID, action)
	actions := [][2]string{{"1", "confirmed"}, {"2", "timeout"}, {"1", "confirmed"}, {"3", "confirmed"}}
	fmt.Println(UsersThatActivelyRequestConfirmationMessages(actions)) // [1 3]
}

// Time: O(n log n), Space: O(n)
func UsersThatActivelyRequestConfirmationMessages(actions [][2]string) []int {
	count := make(map[int]int)
	for _, a := range actions {
		userID := 0
		for _, c := range a[0] {
			userID = userID*10 + int(c-'0')
		}
		if a[1] == "confirmed" {
			count[userID]++
		}
	}

	var result []int
	for uid, c := range count {
		if c >= 2 {
			result = append(result, uid)
		}
	}
	sort.Ints(result)
	return result
}
