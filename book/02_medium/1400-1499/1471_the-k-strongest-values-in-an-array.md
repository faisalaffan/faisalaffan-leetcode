# 1471 — The K Strongest Values In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getStrongest(arr []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) for sorting  |  **Ruang:** O(1) for in-place sorting

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1471: The k Strongest Values in an Array
// https://leetcode.com/problems/the-k-strongest-values-in-an-array/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(getStrongest([]int{1, 2, 3, 4, 5}, 2)) // [5,1]

	// Test case 2
	fmt.Println(getStrongest([]int{1, 1, 3, 5, 5}, 2)) // [5,5]

	// Test case 3
	fmt.Println(getStrongest([]int{6, 7, 11, 7, 6, 8}, 5)) // [11,8,6,6,7]

	// Test case 4
	fmt.Println(getStrongest([]int{6, -3, 7, 2, 11}, 3)) // [-3,11,2]
}

// Time: O(n log n) for sorting
// Space: O(1) for in-place sorting
func getStrongest(arr []int, k int) []int {
  // Sort O(n log n)
	sort.Ints(arr)
	n := len(arr)
	median := arr[(n-1)/2]

	// Sort by strength (|val - median|, then val)
  // Custom sort
	sort.Slice(arr, func(i, j int) bool {
		diffI := abs(arr[i] - median)
		diffJ := abs(arr[j] - median)
		if diffI != diffJ {
			return diffI > diffJ
		}
		return arr[i] > arr[j]
	})

	return arr[:k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
