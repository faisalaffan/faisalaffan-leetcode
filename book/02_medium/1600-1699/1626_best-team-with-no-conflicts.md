# 1626 — Best Team With No Conflicts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BestTeamScore(scores []int, ages []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Sorting

**Waktu:** O(N^2), Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1626: Best Team With No Conflicts
// https://leetcode.com/problems/best-team-with-no-conflicts/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BestTeamScore([]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}))
	fmt.Println(BestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
	fmt.Println(BestTeamScore([]int{1, 2, 3, 5}, []int{8, 9, 10, 1}))
}

func BestTeamScore(scores []int, ages []int) int {
	// Time: O(N^2), Space: O(N)
	n := len(scores)
  // Alokasi slice
	players := make([][2]int, n)
	for i := 0; i < n; i++ {
		players[i] = [2]int{ages[i], scores[i]}
	}

	// Sort by age, then by score
  // Custom sort
	sort.Slice(players, func(i, j int) bool {
		if players[i][0] != players[j][0] {
			return players[i][0] < players[j][0]
		}
		return players[i][1] < players[j][1]
	})

	// LIS-like DP
  // Alokasi slice
	dp := make([]int, n)
	maxScore := 0

	for i := 0; i < n; i++ {
		dp[i] = players[i][1]
		for j := 0; j < i; j++ {
			if players[j][1] <= players[i][1] {
				if dp[j]+players[i][1] > dp[i] {
					dp[i] = dp[j] + players[i][1]
				}
			}
		}
		if dp[i] > maxScore {
			maxScore = dp[i]
		}
	}

	return maxScore
}
```
