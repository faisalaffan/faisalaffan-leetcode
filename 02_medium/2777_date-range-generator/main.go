package main

// LeetCode #2777: Date Range Generator
// https://leetcode.com/problems/date-range-generator/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"time"
)

func DateRangeGenerator(start, end time.Time, step time.Duration) []time.Time {
	dates := make([]time.Time, 0)
	for cur := start; !cur.After(end); cur = cur.Add(step) {
		dates = append(dates, cur)
	}
	return dates
}

func main() {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	dates := DateRangeGenerator(start, end, 24*time.Hour)
	for _, d := range dates {
		fmt.Println(d.Format("2006-01-02"))
	}

	fmt.Println("---")

	start2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	end2 := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	dates2 := DateRangeGenerator(start2, end2, 24*time.Hour)
	for _, d := range dates2 {
		fmt.Println(d.Format("2006-01-02"))
	}
}
