package main

// LeetCode #512: Game Play Analysis II
// https://leetcode.com/problems/game-play-analysis-ii/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisIi() string {
	return "SELECT player_id, device_id FROM Activity WHERE (player_id, event_date) IN (SELECT player_id, MIN(event_date) FROM Activity GROUP BY player_id)"
}

func main() {
	fmt.Println(GamePlayAnalysisIi())
}
