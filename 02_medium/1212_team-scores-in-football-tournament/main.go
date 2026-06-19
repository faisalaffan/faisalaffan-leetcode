package main

import (
	"fmt"
	"sort"
)

// LeetCode #1212: Team Scores in Football Tournament
// https://leetcode.com/problems/team-scores-in-football-tournament/
// Difficulty: Medium [Paid]

// Teams get 3 pts for win, 1 for draw, 0 for loss.
// Calculate total points per team, sorted by points desc then name asc.

// Time: O(n log n)
// Space: O(n)

type matchResult struct {
	host       string
	guest      string
	hostScore  int
	guestScore int
}

type teamScore struct {
	name   string
	points int
}

func calculateTeamScores(teams []string, results []matchResult) []teamScore {
	points := make(map[string]int)

	for _, r := range results {
		if r.hostScore > r.guestScore {
			points[r.host] += 3
		} else if r.hostScore < r.guestScore {
			points[r.guest] += 3
		} else {
			points[r.host] += 1
			points[r.guest] += 1
		}
	}

	scores := make([]teamScore, 0, len(teams))
	for _, t := range teams {
		scores = append(scores, teamScore{t, points[t]})
	}

	sort.Slice(scores, func(i, j int) bool {
		if scores[i].points != scores[j].points {
			return scores[i].points > scores[j].points
		}
		return scores[i].name < scores[j].name
	})

	return scores
}

func main() {
	results := []matchResult{
		{"A", "B", 2, 1},
		{"C", "A", 1, 1},
		{"B", "C", 0, 3},
	}
	scores := calculateTeamScores([]string{"A", "B", "C"}, results)
	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.points)
	}
}
