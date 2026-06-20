# 1893 — Check If All The Integers In A Range Are Covered

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func IsCovered(ranges [][]int, left int, right int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n * range), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1893: Check if All the Integers in a Range Are Covered
// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/
// Difficulty: Easy

import "fmt"

// Time: O(n * range), Space: O(1)
func IsCovered(ranges [][]int, left int, right int) bool {
	covered := make([]bool, 51)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			covered[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !covered[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(IsCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}
```
