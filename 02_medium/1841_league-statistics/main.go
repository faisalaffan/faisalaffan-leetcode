package main

// LeetCode #1841: League Statistics
// https://leetcode.com/problems/league-statistics/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Match struct {
	HomeTeam int
	AwayTeam int
	HomeGoals int
	AwayGoals int
}

type TeamStats struct {
	TeamID   int
	Played   int
	Won      int
	Drawn    int
	Lost     int
	GoalsFor int
	GoalsAgainst int
}

func (ts TeamStats) Points() int {
	return ts.Won*3 + ts.Drawn
}

func leagueStandings(matches []Match) []TeamStats {
	stats := make(map[int]*TeamStats)

	for _, m := range matches {
		if stats[m.HomeTeam] == nil {
			stats[m.HomeTeam] = &TeamStats{TeamID: m.HomeTeam}
		}
		if stats[m.AwayTeam] == nil {
			stats[m.AwayTeam] = &TeamStats{TeamID: m.AwayTeam}
		}

		home := stats[m.HomeTeam]
		away := stats[m.AwayTeam]
		home.Played++
		away.Played++
		home.GoalsFor += m.HomeGoals
		home.GoalsAgainst += m.AwayGoals
		away.GoalsFor += m.AwayGoals
		away.GoalsAgainst += m.HomeGoals

		if m.HomeGoals > m.AwayGoals {
			home.Won++
			away.Lost++
		} else if m.HomeGoals < m.AwayGoals {
			away.Won++
			home.Lost++
		} else {
			home.Drawn++
			away.Drawn++
		}
	}

	result := make([]TeamStats, 0, len(stats))
	for _, ts := range stats {
		result = append(result, *ts)
	}
	sort.Slice(result, func(i, j int) bool {
		pi, pj := result[i].Points(), result[j].Points()
		if pi != pj {
			return pi > pj
		}
		// Goal difference
		gdi := result[i].GoalsFor - result[i].GoalsAgainst
		gdj := result[j].GoalsFor - result[j].GoalsAgainst
		if gdi != gdj {
			return gdi > gdj
		}
		return result[i].GoalsFor > result[j].GoalsFor
	})
	return result
}

func main() {
	matches := []Match{
		{1, 2, 2, 1},
		{1, 3, 1, 1},
		{2, 3, 3, 0},
	}
	standings := leagueStandings(matches)
	for _, s := range standings {
		fmt.Printf("Team %d: %d pts (%dW %dD %dL, %d:%d)\n",
			s.TeamID, s.Points(), s.Won, s.Drawn, s.Lost, s.GoalsFor, s.GoalsAgainst)
	}
}
