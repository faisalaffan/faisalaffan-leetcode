package main

// LeetCode #1369: Get the Second Most Recent Activity
// https://leetcode.com/problems/get-the-second-most-recent-activity/
// Difficulty: Hard [Paid]
//
// Approach: Simulate SQL query in Go.
// From a list of user activities (username, activity, startDate, endDate),
// for each user find the second most recent activity by startDate.
// If a user has only one activity, return that one as the result.
// Results are grouped by username.

import (
	"fmt"
	"sort"
)

type activity struct {
	username  string
	activity  string
	startDate string
	endDate   string
}

func secondMostRecentActivity(activities []activity) []activity {
	byUser := make(map[string][]activity)
	for _, a := range activities {
		byUser[a.username] = append(byUser[a.username], a)
	}

	var result []activity
	for _, acts := range byUser {
		sort.Slice(acts, func(i, j int) bool {
			return acts[i].startDate > acts[j].startDate
		})
		if len(acts) >= 2 {
			result = append(result, acts[1])
		} else {
			result = append(result, acts[0])
		}
	}

	// Sort by username for deterministic output
	sort.Slice(result, func(i, j int) bool {
		return result[i].username < result[j].username
	})
	return result
}

func main() {
	activities := []activity{
		{"Alice", "Travel", "2020-02-12", "2020-02-20"},
		{"Alice", "Dance", "2020-02-21", "2020-02-23"},
		{"Alice", "Travel", "2020-02-24", "2020-02-28"},
		{"Bob", "Travel", "2020-02-11", "2020-02-18"},
		{"Charlie", "Read", "2020-01-01", "2020-01-05"},
		{"Charlie", "Code", "2020-01-10", "2020-01-15"},
		{"Charlie", "Sleep", "2020-01-20", "2020-01-25"},
	}

	result := secondMostRecentActivity(activities)
	for _, a := range result {
		fmt.Printf("%s: %s (%s to %s)\n", a.username, a.activity, a.startDate, a.endDate)
	}
}
