# 3572 — Maximize Ysum By Picking A Triplet Of Distinct Xvalues

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximizeYsumByPickingATripletOfDistinctXvalues(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3572: Maximize Y-Sum by Picking a Triplet of Distinct X-Values
// https://leetcode.com/problems/maximize-ysum-by-picking-a-triplet-of-distinct-xvalues/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	points := [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 5}}
	fmt.Println("Test 1:", MaximizeYsumByPickingATripletOfDistinctXvalues(points))
	// Test case 2
	points2 := [][]int{{1, 5}, {2, 5}, {3, 5}}
	fmt.Println("Test 2:", MaximizeYsumByPickingATripletOfDistinctXvalues(points2))
	// Test case 3
	points3 := [][]int{{1, 10}, {2, 1}, {3, 1}}
	fmt.Println("Test 3:", MaximizeYsumByPickingATripletOfDistinctXvalues(points3))
}

func MaximizeYsumByPickingATripletOfDistinctXvalues(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	// Sort by y descending
  // Custom sort
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] > points[j][1]
	})

  // HashMap: O(1) lookup
	used := make(map[int]bool)
	sum := 0
	count := 0
	for _, p := range points {
		if !used[p[0]] {
			used[p[0]] = true
			sum += p[1]
			count++
			if count == 3 {
				return sum
			}
		}
	}
	return sum
}
```
