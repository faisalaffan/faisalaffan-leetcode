# 0581 — Shortest Unsorted Continuous Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindUnsortedSubarray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #581: Shortest Unsorted Continuous Subarray
// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(FindUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(FindUnsortedSubarray([]int{1}))
}

func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	left := -1
	minRight := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > minRight {
			left = i
		} else {
			minRight = nums[i]
		}
	}

	right := -1
	maxLeft := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < maxLeft {
			right = i
		} else {
			maxLeft = nums[i]
		}
	}

	if right == -1 {
		return 0
	}
	return right - left + 1
}
```
