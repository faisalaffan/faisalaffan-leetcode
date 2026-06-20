# 1301 — Number Of Paths With Max Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func pathsWithMaxScore(board []string) []int
```

> **💡 Hint:** DP from bottom-right to top-left.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1301: Number of Paths with Max Score
// https://leetcode.com/problems/number-of-paths-with-max-score/
// Difficulty: Hard
//
// Approach: DP from bottom-right to top-left.
// For each cell, compute the maximum score achievable from that cell to 'S',
// and the number of distinct paths achieving that score.
// The answer is dp[0][0] = [maxScore, pathCount].
// Result is modulo 1e9+7.
// Board cells: 'E'=start, 'S'=end, 'X'=blocked, digits = score.

import "fmt"

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	const mod = 1_000_000_007

  // Membuat matriks/slice 2D untuk DP
	dpScore := make([][]int, n)
  // Membuat matriks/slice 2D untuk DP
	dpCount := make([][]int, n)
	for i := 0; i < n; i++ {
		dpScore[i] = make([]int, n)
		dpCount[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dpScore[i][j] = -1
		}
	}

	dpScore[n-1][n-1] = 0
	dpCount[n-1][n-1] = 1

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if dpScore[i][j] == -1 {
				continue
			}
			dirs := [][]int{{-1, 0}, {0, -1}, {-1, -1}}
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni < 0 || nj < 0 || board[ni][nj] == 'X' {
					continue
				}
				val := 0
				if board[ni][nj] >= '0' && board[ni][nj] <= '9' {
					val = int(board[ni][nj] - '0')
				}
				newScore := dpScore[i][j] + val
				if newScore > dpScore[ni][nj] {
					dpScore[ni][nj] = newScore
					dpCount[ni][nj] = dpCount[i][j]
				} else if newScore == dpScore[ni][nj] {
					dpCount[ni][nj] = (dpCount[ni][nj] + dpCount[i][j]) % mod
				}
			}
		}
	}

	if dpCount[0][0] == 0 {
		return []int{0, 0}
	}
	return []int{dpScore[0][0], dpCount[0][0]}
}

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E23", "2X2", "12S"})) // [7, 1]
	fmt.Println(pathsWithMaxScore([]string{"E12", "1X1", "21S"})) // [4, 2]
	fmt.Println(pathsWithMaxScore([]string{"E11", "XXX", "11S"})) // [0, 0]
	fmt.Println(pathsWithMaxScore([]string{"E0", "0S"}))          // [0, 3]
}
```
