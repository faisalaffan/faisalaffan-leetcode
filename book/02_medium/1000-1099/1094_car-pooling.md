# 1094 — Car Pooling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func carPooling(trips [][]int, capacity int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + maxLocation)  |  **Ruang:** O(maxLocation)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1094: Car Pooling
// https://leetcode.com/problems/car-pooling/
// Difficulty: Medium
//
// Approach: Difference array (prefix sum)
// Time: O(n + maxLocation)
// Space: O(maxLocation)

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4)) // false
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5)) // true
}

func carPooling(trips [][]int, capacity int) bool {
	maxLoc := 0
	for _, t := range trips {
		if t[2] > maxLoc {
			maxLoc = t[2]
		}
	}

  // Alokasi slice
	diff := make([]int, maxLoc+2)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	current := 0
	for _, d := range diff {
		current += d
		if current > capacity {
			return false
		}
	}

	return true
}
```
