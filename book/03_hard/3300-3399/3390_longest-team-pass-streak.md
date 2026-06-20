# 3390 — Longest Team Pass Streak

## Deskripsi

**Soal:** [3390. Longest Team Pass Streak](https://leetcode.com/problems/longest-team-pass-streak/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3390: Longest Team Pass Streak
// https://leetcode.com/problems/longest-team-pass-streak/
// Difficulty: Hard [Paid]
//
// Sort passes by timestamp. Track consecutive streaks per team.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestTeamPassStreak([][]int{{0, 1, 1}, {1, 1, 2}, {2, 2, 1}}))
}

func LongestTeamPassStreak(passes [][]int) int {
	sort.Slice(passes, func(i, j int) bool {
		return passes[i][0] < passes[j][0]
	})

	type streak struct{ cur, best int }
  // Membuat map untuk pencarian O(1): key → value
	teams := make(map[int]*streak)
	maxStreak := 1

	for _, p := range passes {
		team := p[1]
		if _, ok := teams[team]; !ok {
			teams[team] = &streak{1, 1}
			continue
		}
		teams[team].cur++
		if teams[team].cur > teams[team].best {
			teams[team].best = teams[team].cur
		}
		if teams[team].best > maxStreak {
			maxStreak = teams[team].best
		}
	}

	prevTeam := passes[0][1]
	curStreak := 1
	for i := 1; i < len(passes); i++ {
		if passes[i][1] == prevTeam {
			curStreak++
			if curStreak > maxStreak {
				maxStreak = curStreak
			}
		} else {
			if curStreak > maxStreak {
				maxStreak = curStreak
			}
			curStreak = 1
			prevTeam = passes[i][1]
		}
	}
	if curStreak > maxStreak {
		maxStreak = curStreak
	}
	return maxStreak
}
```
