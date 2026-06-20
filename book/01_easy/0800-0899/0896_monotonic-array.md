# 0896 — Monotonic Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isMonotonic(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Monotonic Stack

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Monotonic Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #896: Monotonic Array
// https://leetcode.com/problems/monotonic-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))   // true
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))   // true
	fmt.Println(isMonotonic([]int{1, 3, 2}))      // false
	fmt.Println(isMonotonic([]int{1, 1, 1}))      // true
}

// isMonotonic checks if the array is monotonic (either non-decreasing or non-increasing).
// Time: O(n). Space: O(1).
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			inc = false
		}
	}
	return inc || dec
}
```
