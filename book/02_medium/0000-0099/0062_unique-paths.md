# 0062 — Unique Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func uniquePaths(m int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #62: Unique Paths
// https://leetcode.com/problems/unique-paths/
// Difficulty: Medium

import "fmt"

func uniquePaths(m int, n int) int {
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(uniquePaths(3, 7)) // 28

	// Test case 2
	fmt.Println(uniquePaths(3, 2)) // 3

	// Test case 3
	fmt.Println(uniquePaths(7, 3)) // 28
}

// Time: O(m*n) | Space: O(m*n)
```
