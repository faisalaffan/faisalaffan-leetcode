# 1277 — Count Square Submatrices With All Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countSquares(matrix [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1277: Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
// Difficulty: Medium

// dp[i][j] = side length of largest square ending at (i,j).
// If grid[i][j]==1: dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
// Sum all dp[i][j] = total squares.

// Time: O(m*n)
// Space: O(m*n)

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				}
				total += dp[i][j]
			}
		}
	}

	return total
}

func main() {
	fmt.Printf("%d (expected: 15)\n",
		countSquares([][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}))

	fmt.Printf("%d (expected: 7)\n",
		countSquares([][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}))
}
```
