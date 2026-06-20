# 3246 — Premier League Table Ranking

## Deskripsi

**Soal:** [3246. Premier League Table Ranking](https://leetcode.com/problems/premier-league-table-ranking/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3246: Premier League Table Ranking
// https://leetcode.com/problems/premier-league-table-ranking/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	stats := []TeamStat{
		{TeamID: 1, TeamName: "City", Wins: 10, Draws: 3, Losses: 2},
		{TeamID: 2, TeamName: "United", Wins: 8, Draws: 5, Losses: 2},
		{TeamID: 3, TeamName: "Liverpool", Wins: 10, Draws: 3, Losses: 2},
	}
	result := PremierLeagueTableRanking(stats)
	for _, r := range result {
		fmt.Println(r)
	}
}

// TeamStat represents a team's statistics.
type TeamStat struct {
	TeamID   int
	TeamName string
	Wins     int
	Draws    int
	Losses   int
}

// TeamRank represents a team's final ranking.
type TeamRank struct {
	TeamID   int
	TeamName string
	Points   int
	Position int
}

// PremierLeagueTableRanking ranks teams by points (3 per win, 1 per draw), with ties sharing the same rank.
// Time: O(n log n). Space: O(n).
func PremierLeagueTableRanking(stats []TeamStat) []TeamRank {
	type team struct {
		id     int
		name   string
		points int
	}
  // Membuat slice untuk menyimpan hasil
	teams := make([]team, len(stats))
	for i, s := range stats {
		teams[i] = team{s.TeamID, s.TeamName, s.Wins*3 + s.Draws}
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points != teams[j].points {
			return teams[i].points > teams[j].points
		}
		return teams[i].name < teams[j].name
	})

  // Membuat slice untuk menyimpan hasil
	result := make([]TeamRank, len(teams))
	for i, t := range teams {
		rank := i + 1
		if i > 0 && t.points == teams[i-1].points {
			rank = result[i-1].Position
		}
		result[i] = TeamRank{t.id, t.name, t.points, rank}
	}
	return result
}
```
