# 0550 — Game Play Analysis Iv

## Deskripsi

**Soal:** [0550. Game Play Analysis Iv](https://leetcode.com/problems/game-play-analysis-iv/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #550: Game Play Analysis IV
// https://leetcode.com/problems/game-play-analysis-iv/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Simulating player activity: {player_id, event_date, games_played}
	activities := [][]int{
		{1, 2016, 5},
		{1, 2017, 10},
		{2, 2016, 15},
		{2, 2018, 20},
		{3, 2020, 30},
		{3, 2021, 25},
	}
	fmt.Printf("%.2f\n", GamePlayAnalysisIv(activities))
}

func GamePlayAnalysisIv(activities [][]int) float64 {
	if len(activities) == 0 {
		return 0.0
	}
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][0] != activities[j][0] {
			return activities[i][0] < activities[j][0]
		}
		return activities[i][1] < activities[j][1]
	})

	// Find first login for each player
  // Membuat map untuk pencarian O(1): key → value
	firstLogin := make(map[int]int)
	for _, act := range activities {
		pid, date := act[0], act[1]
		if _, exists := firstLogin[pid]; !exists {
			firstLogin[pid] = date
		}
	}

	// Count players who logged in the day after their first login
	nextDayPlayers := 0
  // Membuat map untuk pencarian O(1): key → value
	playerSet := make(map[int]bool)
	for _, act := range activities {
		pid, date := act[0], act[1]
		if playerSet[pid] {
			continue
		}
		if date == firstLogin[pid]+1 {
			nextDayPlayers++
			playerSet[pid] = true
		}
	}

	totalPlayers := len(firstLogin)
	if totalPlayers == 0 {
		return 0.0
	}
	return float64(nextDayPlayers) / float64(totalPlayers)
}
```
