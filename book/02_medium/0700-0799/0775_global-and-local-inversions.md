# 0775 — Global And Local Inversions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isIdealPermutation(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #775: Global and Local Inversions
// https://leetcode.com/problems/global-and-local-inversions/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
}

func isIdealPermutation(nums []int) bool {
	maxVal := -1
  // Linear scan O(n)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if maxVal > nums[i+2] {
			return false
		}
	}
	return true
}
```
