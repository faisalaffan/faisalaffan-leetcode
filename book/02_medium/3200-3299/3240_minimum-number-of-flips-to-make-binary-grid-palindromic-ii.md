# 3240 — Minimum Number Of Flips To Make Binary Grid Palindromic Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minFlips(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3240: Minimum Number of Flips to Make Binary Grid Palindromic II
// https://leetcode.com/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-ii/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0

	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			ones := grid[i][j] + grid[i][n-1-j] + grid[m-1-i][j] + grid[m-1-i][n-1-j]
			ans += min(ones, 4-ones)
		}
	}

	mismatchPairs := 0
	onesInMiddle := 0

	if m%2 == 1 {
		mid := m / 2
		for j := 0; j < n/2; j++ {
			if grid[mid][j] != grid[mid][n-1-j] {
				mismatchPairs++
				ans++
			} else if grid[mid][j] == 1 {
				onesInMiddle += 2
			}
		}
	}

	if n%2 == 1 {
		mid := n / 2
		for i := 0; i < m/2; i++ {
			if grid[i][mid] != grid[m-1-i][mid] {
				mismatchPairs++
				ans++
			} else if grid[i][mid] == 1 {
				onesInMiddle += 2
			}
		}
	}

	if m%2 == 1 && n%2 == 1 {
		if grid[m/2][n/2] == 1 {
			ans++
		}
	} else if mismatchPairs == 0 && onesInMiddle%4 != 0 {
		ans += 2
	}

	return ans
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // Expected: 3
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // Expected: 2
}
```
