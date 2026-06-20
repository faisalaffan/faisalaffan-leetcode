# 0741 — Cherry Pickup

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

// LeetCode #741: Cherry Pickup
// https://leetcode.com/problems/cherry-pickup/
// Difficulty: Hard
//
// Algorithm: 3D DP (Two-Person Walk)
// Model as two people walking from (0,0) to (n-1,n-1) simultaneously.
// dp[r1][c1][r2] = max cherries collected when person 1 is at (r1,c1)
// and person 2 is at (r2, c2) where c2 = r1 + c1 - r2 (since steps are equal)
// Handle -1 (thorns) as -infinity.
// Cherries at the same cell are only counted once.

import (
	"fmt"
	"math"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	// dp[r1][c1][r2] where c2 = r1 + c1 - r2
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt32
			}
		}
	}

	// Base case
	dp[0][0][0] = grid[0][0]

	for r1 := 0; r1 < n; r1++ {
		for c1 := 0; c1 < n; c1++ {
			for r2 := 0; r2 < n; r2++ {
				c2 := r1 + c1 - r2
				if c2 < 0 || c2 >= n {
					continue
				}
				if grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}
				if dp[r1][c1][r2] == math.MinInt32 {
					continue
				}

				// Try all 4 combinations of next moves (down, right) for both people
				moves := [][2]int{{1, 0}, {0, 1}}
				for _, m1 := range moves {
					nr1, nc1 := r1+m1[0], c1+m1[1]
					if nr1 >= n || nc1 >= n {
						continue
					}
					if grid[nr1][nc1] == -1 {
						continue
					}
					for _, m2 := range moves {
						nr2, nc2 := r2+m2[0], c2+m2[1]
						if nr2 >= n || nc2 >= n {
							continue
						}
						if grid[nr2][nc2] == -1 {
							continue
						}
						// Add cherries at NEXT positions (count once if same cell)
						add := grid[nr1][nc1]
						if nr1 != nr2 || nc1 != nc2 {
							add += grid[nr2][nc2]
						}
						val := dp[r1][c1][r2] + add
						if val > dp[nr1][nc1][nr2] {
							dp[nr1][nc1][nr2] = val
						}
					}
				}
			}
		}
	}

	result := dp[n-1][n-1][n-1]
	if result < 0 {
		return 0
	}
	return result
}

func main() {
	// Example from problem
	grid1 := [][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}
	result1 := cherryPickup(grid1)
	fmt.Printf("Input: %v\nOutput: %d (expected: 5)\n\n", grid1, result1)

	// Test case 2: 1x1 grid
	grid2 := [][]int{{5}}
	result2 := cherryPickup(grid2)
	fmt.Printf("Input: %v\nOutput: %d (expected: 5)\n\n", grid2, result2)

	// Test case 3: no path
	grid3 := [][]int{
		{0, -1},
		{-1, 1},
	}
	result3 := cherryPickup(grid3)
	fmt.Printf("Input: %v\nOutput: %d (expected: 0)\n\n", grid3, result3)

	// Test case 4: simple 2x2
	grid4 := [][]int{
		{1, 1},
		{1, 1},
	}
	result4 := cherryPickup(grid4)
	fmt.Printf("Input: %v\nOutput: %d (expected: 4)\n\n", grid4, result4)

	// Test case 5: with obstacles
	grid5 := [][]int{
		{0, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}
	result5 := cherryPickup(grid5)
	fmt.Printf("Input (4x4)\nOutput: %d\n\n", result5)

	// Test case 6: all zeros
	grid6 := [][]int{
		{0, 0},
		{0, 0},
	}
	result6 := cherryPickup(grid6)
	fmt.Printf("Input: %v\nOutput: %d (expected: 0)\n\n", grid6, result6)

	// Test case 7: all ones
	grid7 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	result7 := cherryPickup(grid7)
	fmt.Printf("Input: 3x3 all ones\nOutput: %d (expected: 8, start+end shared)\n", result7)
}
```
