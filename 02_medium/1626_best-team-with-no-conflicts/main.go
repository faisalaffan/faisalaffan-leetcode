package main

// LeetCode #1626: Best Team With No Conflicts
// https://leetcode.com/problems/best-team-with-no-conflicts/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BestTeamScore([]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}))
	fmt.Println(BestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
	fmt.Println(BestTeamScore([]int{1, 2, 3, 5}, []int{8, 9, 10, 1}))
}

func BestTeamScore(scores []int, ages []int) int {
	// Time: O(N^2), Space: O(N)
	n := len(scores)
	players := make([][2]int, n)
	for i := 0; i < n; i++ {
		players[i] = [2]int{ages[i], scores[i]}
	}

	// Sort by age, then by score
	sort.Slice(players, func(i, j int) bool {
		if players[i][0] != players[j][0] {
			return players[i][0] < players[j][0]
		}
		return players[i][1] < players[j][1]
	})

	// LIS-like DP
	dp := make([]int, n)
	maxScore := 0

	for i := 0; i < n; i++ {
		dp[i] = players[i][1]
		for j := 0; j < i; j++ {
			if players[j][1] <= players[i][1] {
				if dp[j]+players[i][1] > dp[i] {
					dp[i] = dp[j] + players[i][1]
				}
			}
		}
		if dp[i] > maxScore {
			maxScore = dp[i]
		}
	}

	return maxScore
}
