# 2510 — Check If There Is A Path With Equal Number Of 0S And 1S

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func isThereAPath(grid [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n * (m+n))  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2510: Check if There is a Path With Equal Number of 0's And 1's
// https://leetcode.com/problems/check-if-there-is-a-path-with-equal-number-of-0s-and-1s/
// Difficulty: Medium
// Time: O(m * n * (m+n)) | Space: O(m * n)
// DP: dp[i][j][diff] = reachable with given (ones - zeros) balance.
// Optimize: total cells must be even, and we need diff=0 at end.
// Since grid size is small, we can use 2D boolean DP with reachable sums.

import "fmt"

func main() {
	fmt.Println(isThereAPath([][]int{{0, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // false
	fmt.Println(isThereAPath([][]int{{0, 1}, {1, 0}}))                   // true
}

func isThereAPath(grid [][]int) bool {
	r, c := len(grid), len(grid[0])
	total := r + c - 1
	if total%2 != 0 {
		return false
	}
	target := total / 2 // need this many ones

	// dp[i][j][k] = reachable with k ones
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]bool, r)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]bool, c)
		for j := range dp[i] {
			dp[i][j] = make([]bool, target+1)
		}
	}

	ones := 0
	if grid[0][0] == 1 {
		ones = 1
	}
	if ones <= target {
		dp[0][0][ones] = true
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			for k := 0; k <= target; k++ {
				prev := false
				if i > 0 {
					prev = prev || dp[i-1][j][k]
				}
				if j > 0 {
					prev = prev || dp[i][j-1][k]
				}
				if prev && k+val <= target {
					dp[i][j][k+val] = true
				}
			}
		}
	}
	return dp[r-1][c-1][target]
}
```
