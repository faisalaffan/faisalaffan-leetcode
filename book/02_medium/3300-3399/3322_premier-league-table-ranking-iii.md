# 3322 — Premier League Table Ranking Iii

## Deskripsi

**Soal:** [3322. Premier League Table Ranking Iii](https://leetcode.com/problems/premier-league-table-ranking-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(t log t) Space: O(t)  
**Kompleksitas Ruang:** O(t)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3322: Premier League Table Ranking III
// https://leetcode.com/problems/premier-league-table-ranking-iii/
// Difficulty: Medium
// Time: O(t log t) Space: O(t)

import (
	"fmt"
	"sort"
)

func main() {
	stats := []SeasonStats{
		{1, 1, "City", 38, 28, 6, 4, 80, 20},
		{1, 2, "United", 38, 22, 8, 8, 60, 30},
		{2, 1, "City", 38, 26, 5, 7, 75, 25},
		{2, 2, "Arsenal", 38, 26, 5, 7, 70, 20},
	}
	fmt.Println(premierLeagueRanking(stats))
}

type SeasonStats struct {
	SeasonID       int
	TeamID         int
	TeamName       string
	MatchesPlayed  int
	Wins           int
	Draws          int
	Losses         int
	GoalsFor       int
	GoalsAgainst   int
}

type TeamRank struct {
	SeasonID       int
	TeamID         int
	TeamName       string
	Points         int
	GoalDifference int
}

func premierLeagueRanking(stats []SeasonStats) []TeamRank {
	type teamData struct {
		SeasonID int
		TeamID   int
		Name     string
		Points   int
		GD       int
	}
	var data []teamData
	for _, s := range stats {
		points := s.Wins*3 + s.Draws
		gd := s.GoalsFor - s.GoalsAgainst
		data = append(data, teamData{s.SeasonID, s.TeamID, s.TeamName, points, gd})
	}

	// Group by season and rank
  // Membuat map untuk pencarian O(1): key → value
	seasonTeams := make(map[int][]teamData)
	for _, d := range data {
		seasonTeams[d.SeasonID] = append(seasonTeams[d.SeasonID], d)
	}

	var result []TeamRank
	for _, teams := range seasonTeams {
		sort.Slice(teams, func(i, j int) bool {
			if teams[i].Points != teams[j].Points {
				return teams[i].Points > teams[j].Points
			}
			if teams[i].GD != teams[j].GD {
				return teams[i].GD > teams[j].GD
			}
			return teams[i].Name < teams[j].Name
		})
		for _, t := range teams {
			result = append(result, TeamRank{
				SeasonID: t.SeasonID, TeamID: t.TeamID, TeamName: t.Name,
				Points: t.Points, GoalDifference: t.GD,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].SeasonID != result[j].SeasonID {
			return result[i].SeasonID < result[j].SeasonID
		}
		return result[i].TeamName < result[j].TeamName
	})
	return result
}
```
