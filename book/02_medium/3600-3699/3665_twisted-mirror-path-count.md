# 3665 — Twisted Mirror Path Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func twistedMirrorPathCount(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3665: Twisted Mirror Path Count
// https://leetcode.com/problems/twisted-mirror-path-count/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func twistedMirrorPathCount(grid [][]int) int {
	mod := int(1e9 + 7)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			cur := dp[i][j]

			// Try moving right to (i, j+1)
			if j+1 < n {
				if grid[i][j+1] == 1 {
					// Mirror at (i, j+1): reflect down to (i+1, j+1)
					if i+1 < m {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i][j+1] = (dp[i][j+1] + cur) % mod
				}
			}

			// Try moving down to (i+1, j)
			if i+1 < m {
				if grid[i+1][j] == 1 {
					// Mirror at (i+1, j): reflect right to (i+1, j+1)
					if j+1 < n {
						dp[i+1][j+1] = (dp[i+1][j+1] + cur) % mod
					}
				} else {
					dp[i+1][j] = (dp[i+1][j] + cur) % mod
				}
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 0}, {0, 0}}))
	fmt.Println(twistedMirrorPathCount([][]int{{0, 1}, {1, 0}}))
}
```
