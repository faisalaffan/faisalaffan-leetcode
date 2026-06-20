# 3252 — Premier League Table Ranking Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func premierLeagueRanking(teams [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3252: Premier League Table Ranking II
// https://leetcode.com/problems/premier-league-table-ranking-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type team struct {
	id     int
	points int
	gDiff  int
}

func premierLeagueRanking(teams [][]int) []int {
	var list []team
	for _, t := range teams {
		list = append(list, team{t[0], t[1], t[2]})
	}

  // Custom sort
	sort.Slice(list, func(i, j int) bool {
		if list[i].points != list[j].points {
			return list[i].points > list[j].points
		}
		if list[i].gDiff != list[j].gDiff {
			return list[i].gDiff > list[j].gDiff
		}
		return list[i].id < list[j].id
	})

  // Alokasi slice
	ans := make([]int, len(list))
	for i, t := range list {
		ans[i] = t.id
	}
	return ans
}

func main() {
	fmt.Println(premierLeagueRanking([][]int{{1, 10, 5}, {2, 8, 8}, {3, 10, 3}})) // Expected: [1 3 2]
	fmt.Println(premierLeagueRanking([][]int{{1, 6, 2}, {2, 6, 2}}))              // Expected: [1 2]
}
```
