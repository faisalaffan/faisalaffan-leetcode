package main

// LeetCode #1454: Active Users
// https://leetcode.com/problems/active-users/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Find users who had at least 5 consecutive login days in a given month.

import (
	"fmt"
	"sort"
)

// Login represents a login record.
type Login struct {
	UserID int
	Date   string // format: "YYYY-MM-DD"
}

// findActiveUsers returns user IDs with 5+ consecutive login days.
func findActiveUsers(logins []Login) []int {
	// Group logins by user, deduplicate dates
	userDates := make(map[int]map[string]bool)
	for _, l := range logins {
		if userDates[l.UserID] == nil {
			userDates[l.UserID] = make(map[string]bool)
		}
		userDates[l.UserID][l.Date] = true
	}

	var activeUsers []int
	for userID, dates := range userDates {
		// Convert to sorted slice of ints (days since epoch)
		var dayNums []int
		for d := range dates {
			dayNums = append(dayNums, dateToDays(d))
		}
		sort.Ints(dayNums)

		// Check for 5 consecutive days
		consecutive := 1
		for i := 1; i < len(dayNums); i++ {
			if dayNums[i]-dayNums[i-1] == 1 {
				consecutive++
				if consecutive >= 5 {
					activeUsers = append(activeUsers, userID)
					break
				}
			} else if dayNums[i] != dayNums[i-1] {
				consecutive = 1
			}
		}
	}

	sort.Ints(activeUsers)
	return activeUsers
}

// dateToDays converts a YYYY-MM-DD string to days since epoch.
func dateToDays(date string) int {
	var year, month, day int
	fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)

	// Simple day count (not perfectly accurate for dates before March, but sufficient for consecutive check)
	// Use a month offset table
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Count days from year 0
	total := year * 365
	// Add leap years
	total += (year + 3) / 4
	if year > 0 {
		total -= (year-1)/100 - (year-1)/400
	}

	// Add months for current year
	for m := 1; m < month; m++ {
		total += daysInMonth[m]
		if m == 2 && isLeap(year) {
			total++
		}
	}

	total += day
	return total
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	// Test case 1: single user with 5 consecutive logins
	logins := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{2, "2023-01-01"},
		{2, "2023-01-03"},
		{2, "2023-01-05"},
	}

	active := findActiveUsers(logins)
	fmt.Printf("Test 1 - Active users: %v (expected [1])\n", active)

	// Test case 2: user with logins spanning across month boundary
	logins2 := []Login{
		{1, "2023-01-30"},
		{1, "2023-01-31"},
		{1, "2023-02-01"},
		{1, "2023-02-02"},
		{1, "2023-02-03"},
	}
	active2 := findActiveUsers(logins2)
	fmt.Printf("Test 2 - Active users (month boundary): %v (expected [1])\n", active2)

	// Test case 3: no active users
	logins3 := []Login{
		{1, "2023-06-01"},
		{1, "2023-06-03"},
		{1, "2023-06-05"},
		{2, "2023-06-02"},
	}
	active3 := findActiveUsers(logins3)
	fmt.Printf("Test 3 - Active users (none): %v (expected [])\n", active3)

	// Test case 4: multiple active users
	logins4 := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{2, "2023-02-10"},
		{2, "2023-02-11"},
		{2, "2023-02-12"},
		{2, "2023-02-13"},
		{2, "2023-02-14"},
	}
	active4 := findActiveUsers(logins4)
	fmt.Printf("Test 4 - Active users (multiple): %v (expected [1 2])\n", active4)

	// Test case 5: duplicated dates
	logins5 := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-01"}, // duplicate
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{1, "2023-01-05"}, // duplicate
	}
	active5 := findActiveUsers(logins5)
	fmt.Printf("Test 5 - Active users (duplicates): %v (expected [1])\n", active5)
}
