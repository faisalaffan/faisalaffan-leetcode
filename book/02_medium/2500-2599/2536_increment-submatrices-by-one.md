# 2536 — Increment Submatrices By One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func rangeAddQueries(n int, queries [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2 + q)  |  **Ruang:** O(n^2)


## 💻 Solusi Go

```go
package main

// LeetCode #2536: Increment Submatrices by One
// https://leetcode.com/problems/increment-submatrices-by-one/
// Difficulty: Medium
// Time: O(n^2 + q) | Space: O(n^2)

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
  // Matriks 2D
	diff := make([][]int, n+1)
  // Range loop
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1][c1]++
		diff[r1][c2+1]--
		diff[r2+1][c1]--
		diff[r2+1][c2+1]++
	}

  // Matriks 2D
	mat := make([][]int, n)
  // Range loop
	for i := range mat {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i > 0 {
				diff[i][j] += diff[i-1][j]
			}
			if j > 0 {
				diff[i][j] += diff[i][j-1]
			}
			if i > 0 && j > 0 {
				diff[i][j] -= diff[i-1][j-1]
			}
			mat[i][j] = diff[i][j]
		}
	}
	return mat
}

func main() {
	// Test case 1: n=2, queries=[[0,0,0,0]]
	res1 := rangeAddQueries(2, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 1:", res1)
	// Expected: [[1,0],[0,0]]

	// Test case 2: n=2, queries=[[0,0,1,1]]
	res2 := rangeAddQueries(2, [][]int{{0, 0, 1, 1}})
	fmt.Println("Test 2:", res2)
	// Expected: [[1,1],[1,1]]

	// Test case 3: n=1
	res3 := rangeAddQueries(1, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 3:", res3)
	// Expected: [[1]]
}
```
