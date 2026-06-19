package main

// LeetCode #2720: Popularity Percentage
// https://leetcode.com/problems/popularity-percentage/
// Difficulty: Hard [Paid]
//
// Approach: For each user, count total friends (friendships are bidirectional).
// Popularity percentage = (friend count * 100) / (total users - 1).
// Result is sorted by user ID.

import (
	"fmt"
	"sort"
)

func popularityPercentage(friendships [][]int) [][]int {
	adj := make(map[int]map[int]bool)
	users := make(map[int]bool)

	for _, f := range friendships {
		u, v := f[0], f[1]
		users[u] = true
		users[v] = true
		if adj[u] == nil {
			adj[u] = make(map[int]bool)
		}
		if adj[v] == nil {
			adj[v] = make(map[int]bool)
		}
		adj[u][v] = true
		adj[v][u] = true
	}

	totalUsers := len(users)
	if totalUsers <= 1 {
		if totalUsers == 0 {
			return [][]int{}
		}
		for u := range users {
			return [][]int{{u, 0}}
		}
	}

	type result struct {
		userID     int
		percentage int
	}

	var results []result
	for u := range users {
		pct := len(adj[u]) * 100 / (totalUsers - 1)
		results = append(results, result{u, pct})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].userID < results[j].userID
	})

	out := make([][]int, len(results))
	for i, r := range results {
		out[i] = []int{r.userID, r.percentage}
	}
	return out
}

func main() {
	// Example
	fmt.Println(popularityPercentage([][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}}))
	// Two users
	fmt.Println(popularityPercentage([][]int{{1, 2}}))
	// Linear chain
	fmt.Println(popularityPercentage([][]int{{1, 2}, {2, 3}}))
	// Single user
	fmt.Println(popularityPercentage([][]int{{0, 0}}))
	// Empty
	fmt.Println(popularityPercentage([][]int{}))
}
