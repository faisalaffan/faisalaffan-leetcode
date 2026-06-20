# 1105 — Filling Bookcase Shelves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minHeightShelves(books [][]int, shelfWidth int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1105: Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/
// Difficulty: Medium
//
// Approach: DP. dp[i] = min height to place first i books.
//           Try placing books i-1..j on the same shelf.
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4)) // 6
	fmt.Println(minHeightShelves([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))                               // 4
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
  // Alokasi slice
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 1; i <= n; i++ {
		width := 0
		height := 0
		for j := i; j > 0; j-- {
			width += books[j-1][0]
			if width > shelfWidth {
				break
			}
			if books[j-1][1] > height {
				height = books[j-1][1]
			}
			if dp[j-1]+height < dp[i] {
				dp[i] = dp[j-1] + height
			}
		}
	}

	return dp[n]
}
```
