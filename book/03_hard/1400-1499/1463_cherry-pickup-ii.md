# 1463 — Cherry Pickup Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func cherryPickup(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1463: Cherry Pickup II
// https://leetcode.com/problems/cherry-pickup-ii/
// Difficulty: Hard

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cherryPickup(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// dp[r][c1][c2] = max cherries at row r, robot1 col c1, robot2 col c2
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, rows)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, cols)
		for j := range dp[i] {
			dp[i][j] = make([]int, cols)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][0][cols-1] = grid[0][0] + grid[0][cols-1]

	for r := 1; r < rows; r++ {
		for c1 := 0; c1 < cols; c1++ {
			for c2 := 0; c2 < cols; c2++ {
				prev := -1
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						pc1 := c1 + d1
						pc2 := c2 + d2
						if pc1 >= 0 && pc1 < cols && pc2 >= 0 && pc2 < cols {
							prev = max(prev, dp[r-1][pc1][pc2])
						}
					}
				}
				if prev >= 0 {
					cherries := grid[r][c1] + grid[r][c2]
					if c1 == c2 {
						cherries -= grid[r][c1] // don't double count
					}
					dp[r][c1][c2] = prev + cherries
				}
			}
		}
	}

	ans := 0
	for c1 := 0; c1 < cols; c1++ {
		for c2 := 0; c2 < cols; c2++ {
			ans = max(ans, dp[rows-1][c1][c2])
		}
	}
	return ans
}

func main() {
	// Example: [[3,1,1],[2,5,1],[1,5,5],[2,1,1]] -> 24
	fmt.Println(cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}))
}
```
