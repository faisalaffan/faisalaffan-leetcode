# 0189 — Rotate Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func rotate(nums []int, k int) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #189: Rotate Array
// https://leetcode.com/problems/rotate-array/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	reverse := func(arr []int, l, r int) {
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
	fmt.Println(nums1)

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2)

	nums3 := []int{1}
	rotate(nums3, 0)
	fmt.Println(nums3)
}
```
