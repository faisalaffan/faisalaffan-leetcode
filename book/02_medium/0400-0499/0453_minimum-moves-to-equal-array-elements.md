# 0453 — Minimum Moves To Equal Array Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minMoves(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #453: Minimum Moves to Equal Array Elements
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func minMoves(nums []int) int {
	minVal := math.MaxInt32
	sum := 0
	for _, n := range nums {
		sum += n
		if n < minVal {
			minVal = n
		}
	}
	return sum - minVal*len(nums)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves([]int{1, 2, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 1, 1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1000000000}))
	// Expected: 999999999
}
```
