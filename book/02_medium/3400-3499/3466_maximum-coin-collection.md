# 3466 — Maximum Coin Collection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxCoinCollect(grid [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n) Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3466: Maximum Coin Collection
// https://leetcode.com/problems/maximum-coin-collection/
// Difficulty: Medium [Paid]
// Time: O(m*n) Space: O(m*n)

import "fmt"

func maxCoinCollect(grid [][]int) int64 {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int64, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	dp[0][0] = int64(grid[0][0])
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + int64(grid[0][j])
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + int64(grid[i][0])
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + int64(grid[i][j])
			} else {
				dp[i][j] = dp[i][j-1] + int64(grid[i][j])
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	grid1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(maxCoinCollect(grid1)) // 21

	grid2 := [][]int{{10, 20, 30}}
	fmt.Println(maxCoinCollect(grid2)) // 60

	grid3 := [][]int{{5}}
	fmt.Println(maxCoinCollect(grid3)) // 5
}
```
