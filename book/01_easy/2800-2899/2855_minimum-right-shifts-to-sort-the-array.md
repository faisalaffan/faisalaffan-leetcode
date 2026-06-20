# 2855 — Minimum Right Shifts To Sort The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumRightShiftsToSortTheArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2855: Minimum Right Shifts to Sort the Array
// https://leetcode.com/problems/minimum-right-shifts-to-sort-the-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{1, 3, 5}))
}

func MinimumRightShiftsToSortTheArray(nums []int) int {
	n := len(nums)
	descentIdx := -1
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			if descentIdx != -1 {
				return -1
			}
			descentIdx = i
		}
	}
	if descentIdx == -1 {
		return 0
	}
	if nums[n-1] > nums[0] {
		return -1
	}
	return n - descentIdx - 1
}
```
