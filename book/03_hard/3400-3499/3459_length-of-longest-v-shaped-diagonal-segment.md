# 3459 — Length Of Longest V Shaped Diagonal Segment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lenOfVDiagonal(grid [][]int) int
```

> **💡 Hint:** DP from each cell in all diagonal directions. Track length of

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3459: Length of Longest V-Shaped Diagonal Segment
// https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/
// Difficulty: Hard
//
// Given a grid, find the longest diagonal segment that forms a V shape
// (decreasing then increasing values). A diagonal segment moves in one
// of the four diagonal directions (down-right, down-left, up-right, up-left)
// and must form a V pattern: decreasing for some steps then increasing.
//
// Approach: DP from each cell in all diagonal directions. Track length of
// decreasing and increasing runs from each cell.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lenOfVDiagonal([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Example 2
	fmt.Println(lenOfVDiagonal([][]int{{1, 1}, {1, 1}}))
	// Example 3: single cell
	fmt.Println(lenOfVDiagonal([][]int{{5}}))
	// Edge: 2x3 grid
	fmt.Println(lenOfVDiagonal([][]int{{1, 2}, {3, 4}, {5, 6}}))
}

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// dpDec[i][j][dir] = longest decreasing diagonal starting at (i,j) in direction dir
	// dpInc[i][j][dir] = longest increasing diagonal starting at (i,j) in direction dir
	// dir: 0=down-right, 1=down-left, 2=up-right, 3=up-left
  // Membuat matriks/slice 2D untuk DP
	dpDec := make([][][]int, m)
  // Membuat matriks/slice 2D untuk DP
	dpInc := make([][][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dpDec {
		dpDec[i] = make([][]int, n)
		dpInc[i] = make([][]int, n)
		for j := range dpDec[i] {
			dpDec[i][j] = []int{0, 0, 0, 0}
			dpInc[i][j] = []int{0, 0, 0, 0}
		}
	}

	dirs := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	ans := 1

	// Process cells in reverse diagonal order so dependencies are computed
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for d, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if grid[ni][nj] == grid[i][j]-1 {
						dpDec[i][j][d] = dpDec[ni][nj][d] + 1
					} else {
						dpDec[i][j][d] = 0
					}
					if grid[ni][nj] == grid[i][j]+1 {
						dpInc[i][j][d] = dpInc[ni][nj][d] + 1
					} else {
						dpInc[i][j][d] = 0
					}
				}
			}
		}
	}

	// For each cell, try all pairs of opposite directions
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// V-shape: decreasing in one direction, then increasing in opposite
			for d := 0; d < 4; d++ {
				opp := 3 - d // opposite direction
				dec := dpDec[i][j][d]
				inc := dpInc[i][j][opp]
				if dec > 0 && inc > 0 {
					length := dec + inc + 1
					if length > ans {
						ans = length
					}
				}
				// Just decreasing or just increasing also counts as a longer line
				if dec+1 > ans {
					ans = dec + 1
				}
				if inc+1 > ans {
					ans = inc + 1
				}
			}
		}
	}

	return ans
}
```
