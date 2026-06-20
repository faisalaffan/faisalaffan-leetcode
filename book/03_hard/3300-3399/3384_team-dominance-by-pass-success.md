# 3384 — Team Dominance By Pass Success

## Deskripsi

**Soal:** [3384. Team Dominance By Pass Success](https://leetcode.com/problems/team-dominance-by-pass-success/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3384: Team Dominance by Pass Success
// https://leetcode.com/problems/team-dominance-by-pass-success/
// Difficulty: Hard [Paid]
//
// Aggregate pass stats per team: total passes, successful passes, dominance.

import "fmt"

func main() {
	fmt.Println(TeamDominanceByPassSuccess([][]int{{1, 1, 0}, {1, 1, 1}, {2, 0, 0}}))
}

func TeamDominanceByPassSuccess(passes [][]int) float64 {
	type teamStat struct{ total, succ int }
  // Membuat map untuk pencarian O(1): key → value
	teams := make(map[int]*teamStat)

	for _, p := range passes {
		team, succ := p[0], p[1]
		if _, ok := teams[team]; !ok {
			teams[team] = &teamStat{}
		}
		teams[team].total++
		if succ == 1 {
			teams[team].succ++
		}
	}

	best := 0.0
	for _, s := range teams {
		ratio := float64(s.succ) / float64(s.total)
		if ratio > best {
			best = ratio
		}
	}
	return best
}
```
