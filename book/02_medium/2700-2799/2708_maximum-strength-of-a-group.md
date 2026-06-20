# 2708 — Maximum Strength Of A Group

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxStrength(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2708: Maximum Strength of a Group
// https://leetcode.com/problems/maximum-strength-of-a-group/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}

	positives := []int{}
	negatives := []int{}
	hasZero := false

	for _, v := range nums {
		if v > 0 {
			positives = append(positives, v)
		} else if v < 0 {
			negatives = append(negatives, v)
		} else {
			hasZero = true
		}
	}

  // Sort O(n log n)
	sort.Ints(negatives) // Sort negatives (most negative first)

	ans := int64(1)
	for _, v := range positives {
		ans *= int64(v)
	}

	// Take pairs of negatives (the largest ones, i.e., closest to 0)
	// Sort negatives ascending: [-5, -4, -3, -2, -1]
	// We want to pair them from the least negative end: [-1, -2] * [-3, -4] ...
	if len(negatives)%2 == 1 {
		negatives = negatives[:len(negatives)-1] // Drop the most negative if odd count
	}

	for _, v := range negatives {
		ans *= int64(v)
	}

	if ans == 1 && len(positives) == 0 && len(negatives) == 0 {
		if hasZero {
			return 0
		}
	}

	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxStrength([]int{3, -1, -5, 2, 5, -9}))
	// Expected: 1350

	// Test case 2
	fmt.Println("Test 2:", maxStrength([]int{-4, -5, -6}))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", maxStrength([]int{0, -1}))
	// Expected: 0
}
```
