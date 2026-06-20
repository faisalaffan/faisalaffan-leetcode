# 0969 — Pancake Sorting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func pancakeSort(arr []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #969: Pancake Sorting
// https://leetcode.com/problems/pancake-sorting/
// Difficulty: Medium

import "fmt"

// Time: O(n^2) | Space: O(n)
func pancakeSort(arr []int) []int {
  // Alokasi slice
	ans := make([]int, 0)
	n := len(arr)

	for i := n; i > 1; i-- {
		// Find index of max in arr[:i]
		maxIdx := 0
		for j := 1; j < i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx == i-1 {
			continue
		}
		// Flip to bring max to front
		if maxIdx > 0 {
			reverse(arr, maxIdx)
			ans = append(ans, maxIdx+1)
		}
		// Flip to put max at correct position
		reverse(arr, i-1)
		ans = append(ans, i)
	}

	return ans
}

func reverse(arr []int, end int) {
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
	fmt.Println(pancakeSort([]int{1, 2, 3}))
}
```
