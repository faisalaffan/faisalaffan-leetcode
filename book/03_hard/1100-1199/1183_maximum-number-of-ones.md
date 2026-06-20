# 1183 — Maximum Number Of Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1183: Maximum Number of Ones
// https://leetcode.com/problems/maximum-number-of-ones/
// Difficulty: Hard [Paid]
//
// Given a width x height matrix, we can place at most maxOnes ones in any
// sideLength x sideLength submatrix. Find the maximum total number of ones
// we can place in the entire matrix.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: width=3, height=3, sideLength=2, maxOnes=1 => 4
	fmt.Println(maximumNumberOfOnes(3, 3, 2, 1)) // 4

	// Example 2: width=3, height=3, sideLength=2, maxOnes=2 => 6
	fmt.Println(maximumNumberOfOnes(3, 3, 2, 2)) // 6

	// Single cell
	fmt.Println(maximumNumberOfOnes(1, 1, 1, 1)) // 1
	fmt.Println(maximumNumberOfOnes(1, 1, 1, 0)) // 0

	// Full coverage - sideLength equals dimension
	fmt.Println(maximumNumberOfOnes(10, 10, 10, 5)) // 5

	// Rectangular
	fmt.Println(maximumNumberOfOnes(4, 5, 2, 3)) // 12
}

// maximumNumberOfOnes computes the maximum total ones that can be placed in a
// width x height matrix such that any sideLength x sideLength submatrix has at
// most maxOnes ones.
//
// Key insight: The constraint is uniform across all overlapping windows.
// Consider the matrix tiled with sideLength x sideLength blocks. Each position
// (i,j) in a block corresponds to positions (i + a*sideLength, j + b*sideLength)
// in the full matrix. The number of submatrices (windows) covering a cell
// depends on its position within the block pattern.
//
// We count, for each of the sideLength^2 positions within the pattern, how many
// times that relative position appears in the full matrix (i.e., its "coverage"
// count). We then sort these counts descending and place ones at the maxOnes
// positions with the highest coverage.
func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
	// For each position (i, j) in a sideLength x sideLength block, compute how
	// many times it is covered (i.e., appears in the full matrix).
  // Alokasi slice
	counts := make([]int, 0, sideLength*sideLength)

	for i := 0; i < sideLength; i++ {
		for j := 0; j < sideLength; j++ {
			// Number of times this relative position repeats across width
			// The positions are i, i+sideLength, i+2*sideLength, ...
			// Up to width (inclusive of the start) for horizontal,
			// or height for vertical.
			cw := (width - i - 1) / sideLength + 1
			ch := (height - j - 1) / sideLength + 1
			counts = append(counts, cw*ch)
		}
	}

	// Sort descending so we pick the positions with highest coverage first
  // Custom sort
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	total := 0
	for k := 0; k < maxOnes && k < len(counts); k++ {
		total += counts[k]
	}
	return total
}
```
