# 1725 — Number Of Rectangles That Can Form The Largest Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CountGoodRectangles(rectangles [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1725: Number of Rectangles That Can Form the Largest Square
// https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	count := 0
	for _, rect := range rectangles {
		side := rect[0]
		if rect[1] < side {
			side = rect[1]
		}
		if side > maxLen {
			maxLen = side
			count = 1
		} else if side == maxLen {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
	fmt.Println(CountGoodRectangles([][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}))
}
```
