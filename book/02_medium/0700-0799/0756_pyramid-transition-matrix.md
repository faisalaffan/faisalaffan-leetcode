# 0756 — Pyramid Transition Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func pyramidTransition(bottom string, allowed []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(7^b) worst case where b is number of blocks  |  **Ruang:** O(7^b)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #756: Pyramid Transition Matrix
// https://leetcode.com/problems/pyramid-transition-matrix/
// Difficulty: Medium
// Time: O(7^b) worst case where b is number of blocks
// Space: O(7^b)

import "fmt"

func main() {
	fmt.Println(pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"}))
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
}

func pyramidTransition(bottom string, allowed []string) bool {
  // HashMap: O(1) lookup
	memo := make(map[string]bool)
  // HashMap: O(1) lookup
	patterns := make(map[string][]byte)

	for _, a := range allowed {
		key := a[:2]
		patterns[key] = append(patterns[key], a[2])
	}

	var dfs func(row string, next string, idx int) bool
	dfs = func(row string, next string, idx int) bool {
		if len(row) == 1 {
			return true
		}

		key := row + "#" + next
		if val, ok := memo[key]; ok {
			return val
		}

		if idx == len(row)-1 {
			if dfs(next, "", 0) {
				memo[key] = true
				return true
			}
			memo[key] = false
			return false
		}

		chars := patterns[row[idx:idx+2]]
		for _, c := range chars {
			if dfs(row, next+string(c), idx+1) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(bottom, "", 0)
}
```
