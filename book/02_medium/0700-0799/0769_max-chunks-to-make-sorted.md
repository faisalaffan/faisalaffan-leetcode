# 0769 — Max Chunks To Make Sorted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxChunksToSorted(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #769: Max Chunks To Make Sorted
// https://leetcode.com/problems/max-chunks-to-make-sorted/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}

func maxChunksToSorted(arr []int) int {
	count := 0
	maxVal := 0

	for i, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if maxVal == i {
			count++
		}
	}

	return count
}
```
