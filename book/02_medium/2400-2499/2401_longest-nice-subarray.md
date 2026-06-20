# 2401 — Longest Nice Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestNiceSubarray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2401: Longest Nice Subarray
// https://leetcode.com/problems/longest-nice-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: maintain OR of current window. If new element conflicts, shrink left.

import "fmt"

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))  // 1
}

func longestNiceSubarray(nums []int) int {
	left, orMask, ans := 0, 0, 0
	for right, v := range nums {
		for orMask&v != 0 {
			orMask ^= nums[left]
			left++
		}
		orMask |= v
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
```
