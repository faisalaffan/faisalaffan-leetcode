# 1151 — Minimum Swaps To Group All 1S Together

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minSwaps(data []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sliding Window** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1151: Minimum Swaps to Group All 1's Together
// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together/
// Difficulty: Medium [Paid]

// Use sliding window of size = total number of 1s.
// Minimum swaps = window with maximum 1s (least 0s to swap out).

// Time: O(n)
// Space: O(1)

func minSwaps(data []int) int {
	totalOnes := 0
	for _, v := range data {
		if v == 1 {
			totalOnes++
		}
	}
	if totalOnes <= 1 {
		return 0
	}

	currOnes := 0
	for i := 0; i < totalOnes; i++ {
		if data[i] == 1 {
			currOnes++
		}
	}

	maxOnes := currOnes
	for i := totalOnes; i < len(data); i++ {
		if data[i] == 1 {
			currOnes++
		}
		if data[i-totalOnes] == 1 {
			currOnes--
		}
		if currOnes > maxOnes {
			maxOnes = currOnes
		}
	}

	return totalOnes - maxOnes
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minSwaps([]int{1, 0, 1, 0, 1}))
	fmt.Printf("%d (expected: 0)\n", minSwaps([]int{0, 0, 0, 1, 0}))
	fmt.Printf("%d (expected: 2)\n", minSwaps([]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}))
}
```
