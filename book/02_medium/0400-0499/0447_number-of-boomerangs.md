# 0447 — Number Of Boomerangs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numberOfBoomerangs(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #447: Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	total := 0

  // Linear scan O(n)
	for i := 0; i < len(points); i++ {
  // HashMap: O(1) lookup
		distCount := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := dx*dx + dy*dy
			distCount[dist]++
		}
		for _, count := range distCount {
			total += count * (count - 1)
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numberOfBoomerangs([][]int{{0, 0}}))
	// Expected: 0
}
```
