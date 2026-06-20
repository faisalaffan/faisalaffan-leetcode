# 1752 — Check If Array Is Sorted And Rotated

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Check(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1752: Check if Array Is Sorted and Rotated
// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func Check(nums []int) bool {
	drops := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			drops++
		}
	}
	return drops <= 1
}

func main() {
	fmt.Println(Check([]int{3, 4, 5, 1, 2}))
	fmt.Println(Check([]int{2, 1, 3, 4}))
	fmt.Println(Check([]int{1, 2, 3}))
}
```
