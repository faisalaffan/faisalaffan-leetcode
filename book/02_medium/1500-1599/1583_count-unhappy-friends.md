# 1583 — Count Unhappy Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func UnhappyFriends(n int, preferences [][]int, pairs [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N^2), Space: O(N^2)  |  **Ruang:** O(N^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1583: Count Unhappy Friends
// https://leetcode.com/problems/count-unhappy-friends/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(UnhappyFriends(4, [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}, [][]int{{0, 1}, {2, 3}}))
	fmt.Println(UnhappyFriends(2, [][]int{{1}, {0}}, [][]int{{0, 1}}))
	fmt.Println(UnhappyFriends(4, [][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}}, [][]int{{0, 1}, {2, 3}}))
}

func UnhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	// Time: O(N^2), Space: O(N^2)
	// Build preference rank matrix: rank[i][j] = how much i prefers j
  // Matriks 2D
	rank := make([][]int, n)
	for i := 0; i < n; i++ {
		rank[i] = make([]int, n)
		for pos, j := range preferences[i] {
			rank[i][j] = pos
		}
	}

	// Map partner
  // HashMap: O(1) lookup
	partner := make(map[int]int)
	for _, p := range pairs {
		partner[p[0]] = p[1]
		partner[p[1]] = p[0]
	}

	unhappy := 0

	for x := 0; x < n; x++ {
		y := partner[x]
		for _, u := range preferences[x] {
			if u == y {
				break
			}
			// x prefers u over y
			v := partner[u]
			// Is u unhappy? Check if u prefers x over v
			if rank[u][x] < rank[u][v] {
				unhappy++
				break
			}
		}
	}

	return unhappy
}
```
