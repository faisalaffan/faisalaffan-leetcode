# 3393 — Count Paths With The Given Xor Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countPathsWithXorValue(grid [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n*U) Space: O(m*n*U) where U = maxXOR value  |  **Ruang:** O(m*n*U) where U = maxXOR value


## 💻 Solusi Go

```go
package main

// LeetCode #3393: Count Paths With the Given XOR Value
// https://leetcode.com/problems/count-paths-with-the-given-xor-value/
// Difficulty: Medium
// Time: O(m*n*U) Space: O(m*n*U) where U = maxXOR value

import "fmt"

func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	u := 1
	for _, row := range grid {
		for _, val := range row {
			for u <= val {
				u <<= 1
			}
		}
	}
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
  // Matriks 2D
	f := make([][][]int, m+1)
  // Range loop
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, u)
		}
	}

	f[0][1][0] = 1
	for i, row := range grid {
		for j, val := range row {
			for x := 0; x < u; x++ {
				f[i+1][j+1][x] = (f[i+1][j][x^val] + f[i][j+1][x^val]) % mod
			}
		}
	}
	return f[m][n][k]
}

func main() {
	fmt.Println(countPathsWithXorValue([][]int{{2, 1, 5}, {7, 10, 0}, {12, 6, 4}}, 11)) // 3
	fmt.Println(countPathsWithXorValue([][]int{{1, 3, 3, 3}, {0, 3, 3, 2}, {3, 0, 1, 1}}, 2)) // 4
}
```
