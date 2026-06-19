package main

// LeetCode #1132: Reported Posts II
// https://leetcode.com/problems/reported-posts-ii/
// Difficulty: Medium
//
// Approach: Calculate average daily spam removal percentage
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// (date, post_id, action_type) where action_type: 0=report, 1=remove
	// Day 1: post 1 reported, post 2 reported
	// Day 2: post 1 removed (counts for day 1)
	// Day 3: post 3 reported
	// Day 1: remove rate = 1/2 = 50%. Day 3: remove rate = 0/1 = 0%. Avg = 25%
	actions := [][3]int{
		{1, 1, 0},
		{1, 2, 0},
		{2, 1, 1},
		{3, 3, 0},
	}
	fmt.Println(reportedPostsIi(actions))
}

func reportedPostsIi(actions [][3]int) float64 {
	// First pass: identify reported posts per date
	reported := make(map[int]map[int]bool)
	removedPosts := make(map[int]bool) // posts that were eventually removed

	for _, a := range actions {
		date, postID, actionType := a[0], a[1], a[2]
		if actionType == 0 { // report
			if reported[date] == nil {
				reported[date] = make(map[int]bool)
			}
			reported[date][postID] = true
		} else { // remove
			removedPosts[postID] = true
		}
	}

	totalPct := 0.0
	dayCount := 0

	for _, posts := range reported {
		reportCount := len(posts)
		if reportCount > 0 {
			removeCount := 0
			for postID := range posts {
				if removedPosts[postID] {
					removeCount++
				}
			}
			totalPct += float64(removeCount) / float64(reportCount) * 100
			dayCount++
		}
	}

	if dayCount == 0 {
		return 0
	}
	return totalPct / float64(dayCount)
}
