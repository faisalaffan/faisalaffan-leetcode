# 2848 — Points That Intersect With Cars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func PointsThatIntersectWithCars(nums [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * range)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2848: Points That Intersect With Cars
// https://leetcode.com/problems/points-that-intersect-with-cars/
// Difficulty: Easy
// Time: O(n * range) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(PointsThatIntersectWithCars([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(PointsThatIntersectWithCars([][]int{{1, 3}, {5, 8}}))
}

func PointsThatIntersectWithCars(nums [][]int) int {
  // Alokasi slice
	points := make([]bool, 101)
	for _, car := range nums {
		for p := car[0]; p <= car[1]; p++ {
			points[p] = true
		}
	}
	count := 0
	for _, v := range points {
		if v {
			count++
		}
	}
	return count
}
```
