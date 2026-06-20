# 3882 — Minimum Xor Path In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumXorPathInAGrid(grid [][]int) int
```

> **💡 Hint:** DP tracking reachable XOR values at each cell. Only right/down moves.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(M * N * 2^B)  
**Kompleksitas Ruang:** O(N * 2^B) where B = 11 (grid values < 1024)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3882: Minimum XOR Path in a Grid
// https://leetcode.com/problems/minimum-xor-path-in-a-grid/
// Difficulty: Medium
// Time: O(M * N * 2^B) | Space: O(N * 2^B) where B = 11 (grid values < 1024)
// Approach: DP tracking reachable XOR values at each cell. Only right/down moves.

import (
	"fmt"
	"math"
)

func MinimumXorPathInAGrid(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
  // Edge case: input kosong — langsung return
	if n == 0 {
		return -1
	}

	maxXor := 2048 // 2^11 since grid[i][j] <= 1023

	// Use bitset (boolean array) for each cell
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]bool, n)
	for j := 0; j < n; j++ {
		dp[j] = make([]bool, maxXor)
	}

	dp[0][grid[0][0]] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cur := make([]bool, maxXor)
			if i > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			if j > 0 {
				for x := 0; x < maxXor; x++ {
					if dp[j-1][x] {
						cur[x^grid[i][j]] = true
					}
				}
			}
			dp[j] = cur
		}
	}

	ans := math.MaxInt32
	for x := 0; x < maxXor; x++ {
		if dp[n-1][x] && x < ans {
			ans = x
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumXorPathInAGrid([][]int{{1, 2}, {3, 4}})) // Expected: 6

	// Example 2
	fmt.Println(MinimumXorPathInAGrid([][]int{{6, 7}, {5, 8}})) // Expected: 9

	// Example 3
	fmt.Println(MinimumXorPathInAGrid([][]int{{2, 7, 5}})) // Expected: 0
}
```
