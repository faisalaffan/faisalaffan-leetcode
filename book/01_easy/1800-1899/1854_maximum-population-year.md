# 1854 — Maximum Population Year

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximumPopulation(logs [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + range), Space: O(range)  |  **Ruang:** O(range)


## 💻 Solusi Go

```go
package main

// LeetCode #1854: Maximum Population Year
// https://leetcode.com/problems/maximum-population-year/
// Difficulty: Easy

import "fmt"

// Time: O(n + range), Space: O(range)
func MaximumPopulation(logs [][]int) int {
  // Alokasi slice
	delta := make([]int, 101) // 1950 to 2050
	for _, log := range logs {
		delta[log[0]-1950]++
		delta[log[1]-1950]--
	}
	maxPop := 0
	currentPop := 0
	bestYear := 1950
	for i := 0; i < 101; i++ {
		currentPop += delta[i]
		if currentPop > maxPop {
			maxPop = currentPop
			bestYear = 1950 + i
		}
	}
	return bestYear
}

func main() {
	fmt.Println(MaximumPopulation([][]int{{1993, 1999}, {2000, 2010}}))
	fmt.Println(MaximumPopulation([][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}))
}
```
