package main

// LeetCode #511: Game Play Analysis I
// https://leetcode.com/problems/game-play-analysis-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisI() string {
	return "SELECT player_id, MIN(event_date) AS first_login FROM Activity GROUP BY player_id"
}

func main() {
	fmt.Println(GamePlayAnalysisI())
}
