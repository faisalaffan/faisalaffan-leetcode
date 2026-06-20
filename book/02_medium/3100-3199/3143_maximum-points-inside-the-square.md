# 3143 — Maximum Points Inside The Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxPointsInsideSquare(points [][]int, s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3143: Maximum Points Inside the Square
// https://leetcode.com/problems/maximum-points-inside-the-square/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxPointsInsideSquare(points [][]int, s string) int {
	type item struct {
		dist int
		tag  byte
	}
	n := len(points)
	arr := make([]item, n)
	for i, p := range points {
		d := max(abs(p[0]), abs(p[1]))
		arr[i] = item{d, s[i]}
	}

  // Custom sort
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].dist < arr[j].dist
	})

  // HashMap: O(1) lookup
	seen := make(map[byte]bool)
	ans := 0
	i := 0
	for i < n {
		j := i
		for j < n && arr[j].dist == arr[i].dist {
			if seen[arr[j].tag] {
				return ans
			}
			j++
		}
		for k := i; k < j; k++ {
			seen[arr[k].tag] = true
		}
		ans = len(seen)
		i = j
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maxPointsInsideSquare([][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca"))
}
```
