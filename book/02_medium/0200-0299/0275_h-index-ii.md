# 0275 — H Index Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func hIndex(citations []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #275: H-Index II
// https://leetcode.com/problems/h-index-ii/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1

  // Binary search loop
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] >= n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return n - left
}

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{1, 2, 100}))
	fmt.Println(hIndex([]int{0}))
}
```
