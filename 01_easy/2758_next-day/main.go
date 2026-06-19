package main

// LeetCode #2758: Next Day
// https://leetcode.com/problems/next-day/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS Date problem adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(NextDay("2024-03-01"))
	fmt.Println(NextDay("2024-12-31"))
}

func NextDay(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return t.AddDate(0, 0, 1).Format("2006-01-02")
}
