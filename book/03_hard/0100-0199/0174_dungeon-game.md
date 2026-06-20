# 0174 — Dungeon Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateMinimumHP(dungeon [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #174: Dungeon Game
// https://leetcode.com/problems/dungeon-game/
// Difficulty: Hard

import (
	"fmt"
)

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	m, n := len(dungeon), len(dungeon[0])

	// dp[i][j] = minimum health needed to reach bottom-right from (i, j)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				// Bottom-right cell
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				// Last row, can only go right
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				// Last column, can only go down
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				// Can go right or down, take minimum health path
				minNext := min(dp[i][j+1], dp[i+1][j])
				dp[i][j] = max(1, minNext-dungeon[i][j])
			}
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	result := calculateMinimumHP(dungeon)
	expected := 7

	fmt.Printf("calculateMinimumHP(%v) = %d\n", dungeon, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
