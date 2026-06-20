# 2549 — Count Distinct Numbers On Board

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CountDistinctNumbersOnBoard(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2549: Count Distinct Numbers on Board
// https://leetcode.com/problems/count-distinct-numbers-on-board/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountDistinctNumbersOnBoard(5)) // 4
	fmt.Println(CountDistinctNumbersOnBoard(2)) // 1
	fmt.Println(CountDistinctNumbersOnBoard(1)) // 1
}

func CountDistinctNumbersOnBoard(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
```
