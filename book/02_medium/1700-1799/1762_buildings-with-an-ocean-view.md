# 1762 — Buildings With An Ocean View

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findBuildings(heights []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1) excluding output  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #1762: Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) excluding output

import "fmt"

func findBuildings(heights []int) []int {
  // Alokasi slice
	result := make([]int, 0)
	maxHeight := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append([]int{i}, result...)
			maxHeight = heights[i]
		}
	}
	return result
}

func main() {
	fmt.Println(findBuildings([]int{4, 2, 3, 1}))    // Expected: [0, 2, 3]
	fmt.Println(findBuildings([]int{4, 3, 2, 1}))    // Expected: [0, 1, 2, 3]
	fmt.Println(findBuildings([]int{1, 3, 2, 4}))    // Expected: [3]
}
```
