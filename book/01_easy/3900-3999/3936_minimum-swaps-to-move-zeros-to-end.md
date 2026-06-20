# 3936 — Minimum Swaps To Move Zeros To End

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumSwapsToMoveZerosToEnd(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3936: Minimum Swaps to Move Zeros to End
// https://leetcode.com/problems/minimum-swaps-to-move-zeros-to-end/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{0, 1, 0, 3, 12}))
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{0, 1, 0, 2}))
	fmt.Println(MinimumSwapsToMoveZerosToEnd([]int{1, 2, 0}))
}

// Time: O(n)
// Space: O(1)
func MinimumSwapsToMoveZerosToEnd(nums []int) int {
	zeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		}
	}
	boundary := len(nums) - zeroCount
	swaps := 0
	for i := 0; i < boundary; i++ {
		if nums[i] == 0 {
			swaps++
		}
	}
	return swaps
}
```
