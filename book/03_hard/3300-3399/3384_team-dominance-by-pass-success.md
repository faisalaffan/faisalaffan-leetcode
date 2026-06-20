# 3384 — Team Dominance By Pass Success

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func TeamDominanceByPassSuccess(passes [][]int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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
