# 1733 — Minimum Number Of People To Teach

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minimumTeachings(n int, languages [][]int, friendships [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m + f) where n = user count, m = avg languages per user, f = friend pairs  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1733: Minimum Number of People to Teach
// https://leetcode.com/problems/minimum-number-of-people-to-teach/
// Difficulty: Medium
// Time: O(n * m + f) where n = user count, m = avg languages per user, f = friend pairs

import "fmt"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
  // Alokasi slice
	langSet := make([]map[int]bool, len(languages))
	for i, langs := range languages {
		langSet[i] = make(map[int]bool)
		for _, l := range langs {
			langSet[i][l] = true
		}
	}

	// Find users who cannot communicate
	cannotCommunicate := make([]bool, len(languages))
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1
		canCommunicate := false
		for l := range langSet[u] {
			if langSet[v][l] {
				canCommunicate = true
				break
			}
		}
		if !canCommunicate {
			cannotCommunicate[u] = true
			cannotCommunicate[v] = true
		}
	}

	// Find the most common language among users who cannot communicate
  // HashMap: O(1) lookup
	langCount := make(map[int]int)
	for i, cn := range cannotCommunicate {
		if cn {
			for l := range langSet[i] {
				langCount[l]++
			}
		}
	}

	// Count users who don't know the most common language
	maxLang := 0
	for _, c := range langCount {
		if c > maxLang {
			maxLang = c
		}
	}

	totalCannot := 0
	for _, cn := range cannotCommunicate {
		if cn {
			totalCannot++
		}
	}

	return totalCannot - maxLang
}

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}})) // Expected: 1
	fmt.Println(minimumTeachings(3, [][]int{{2}, {1, 3}, {1, 2}, {3}}, [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}})) // Expected: 2
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}}, [][]int{{1, 2}})) // Expected: 1
}
```
