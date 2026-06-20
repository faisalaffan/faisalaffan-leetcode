# 2267 — Check If There Is A Valid Parentheses String Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func hasValidPath(grid [][]byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2267: Check if There Is a Valid Parentheses String Path
// https://leetcode.com/problems/check-if-there-is-a-valid-parentheses-string-path/
// Difficulty: Hard
//
// Given a grid of '(' and ')', find a path from (0,0) to (m-1,n-1)
// moving only down or right, such that the concatenated string is a
// valid parentheses string. A valid parentheses string has:
//   - equal number of '(' and ')'
//   - at any prefix, #('(') >= #(')')

import (
	"fmt"
)

// hasValidPath returns true if a valid path exists.
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	// If start or end is invalid immediately
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}
	// Path length must be even for balanced parens
	if (m+n-1)%2 != 0 {
		return false
	}

	// DP with set of possible open counts at each cell
	// dp[i][j] = set of possible open parenthesis counts when reaching (i,j)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]map[int]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]map[int]bool, n)
	}

	// Initialize
	dp[0][0] = map[int]bool{1: true} // '(' at (0,0) -> open count = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dp[i][j] = make(map[int]bool)

  // Membuat map (HashMap) — pencarian O(1)
			possibleOpen := make(map[int]bool)
			// from top
			if i > 0 {
				for cnt := range dp[i-1][j] {
					possibleOpen[cnt] = true
				}
			}
			// from left
			if j > 0 {
				for cnt := range dp[i][j-1] {
					possibleOpen[cnt] = true
				}
			}

			for cnt := range possibleOpen {
				if grid[i][j] == '(' {
					dp[i][j][cnt+1] = true
				} else { // ')'
					if cnt-1 >= 0 {
						dp[i][j][cnt-1] = true
					}
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}

func main() {
	// Example 1
	grid1 := [][]byte{
		{'(', '(', '('},
		{')', '(', ')'},
		{'(', '(', ')'},
		{'(', '(', ')'},
	}
	fmt.Println(hasValidPath(grid1)) // Expected: true

	// Example 2
	grid2 := [][]byte{
		{')', ')'},
		{'(', '('},
	}
	fmt.Println(hasValidPath(grid2)) // Expected: false
}
```
