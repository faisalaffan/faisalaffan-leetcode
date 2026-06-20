# 0026 — Remove Duplicates From Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RemoveDuplicates(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #26: Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveDuplicates(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{1, 1, 2}
	fmt.Println(RemoveDuplicates(n1), n1)
	n2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicates(n2), n2)
}
```
