# 2906 — Construct Product Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func constructProductMatrix(grid [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2906: Construct Product Matrix
// https://leetcode.com/problems/construct-product-matrix/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructProductMatrix([][]int{{1, 2}, {3, 4}}))
	fmt.Println(constructProductMatrix([][]int{{2, 3, 4}, {5, 6, 7}}))
}

func constructProductMatrix(grid [][]int) [][]int {
	const mod int = 12345
	n, m := len(grid), len(grid[0])
  // Matriks 2D
	p := make([][]int, n)
  // Range loop
	for i := range p {
		p[i] = make([]int, m)
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = p[i][j] * pre % mod
			pre = pre * grid[i][j] % mod
		}
	}
	return p
}
```
