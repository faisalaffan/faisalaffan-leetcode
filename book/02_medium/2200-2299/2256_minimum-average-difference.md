# 2256 — Minimum Average Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumAverageDifference(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2256: Minimum Average Difference
// https://leetcode.com/problems/minimum-average-difference/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	var prefix int64 = 0
	minDiff := int64(1 << 60)
	minIdx := 0

	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		left := prefix / int64(i+1)
		var right int64 = 0
		if i < n-1 {
			right = (total - prefix) / int64(n-i-1)
		}
		diff := left - right
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
			minIdx = i
		}
	}
	return minIdx
}

func main() {
	// Test case 1
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumAverageDifference([]int{0}))
	// Expected: 0

	// Test case 3
	fmt.Println(minimumAverageDifference([]int{0, 1, 0, 1, 0, 1}))
	// Expected: 1
}
```
