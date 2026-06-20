# 3379 — Transformed Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TransformedArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3379: Transformed Array
// https://leetcode.com/problems/transformed-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(TransformedArray([]int{-1, 4, -1}))
}

// TransformedArray constructs a new array where result[i] = nums[(i + nums[i]) mod n], handling negative wrap-around.
// Time: O(n). Space: O(n).
func TransformedArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)
	for i, val := range nums {
		idx := (i + val) % n
		if idx < 0 {
			idx += n
		}
		result[i] = nums[idx]
	}
	return result
}
```
