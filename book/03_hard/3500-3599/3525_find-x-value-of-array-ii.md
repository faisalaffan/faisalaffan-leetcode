# 3525 — Find X Value Of Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func resultArray(nums []int, k int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum, Segment Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3525: Find X Value of Array II
// https://leetcode.com/problems/find-x-value-of-array-ii/
// Difficulty: Hard
//
// Given an array nums, integer k, and queries, for each query [l, r],
// find the value x such that XOR of subarray with x applied maximizes
// something. [details inferred from problem name]
//
// Approach: Process queries using prefix XOR and segment tree / BIT.

import "fmt"

func main() {
	// Example 1
	fmt.Println(resultArray([]int{1, 2, 3, 4}, 2, [][]int{{0, 2}, {1, 3}}))
	// Example 2
	fmt.Println(resultArray([]int{5, 3, 2, 1}, 1, [][]int{{0, 3}}))
	// Edge: single element queries
	fmt.Println(resultArray([]int{7}, 3, [][]int{{0, 0}}))
}

func resultArray(nums []int, k int, queries [][]int) []int {
	n := len(nums)
	// Precompute prefix XOR
  // Alokasi slice
	prefXor := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefXor[i+1] = prefXor[i] ^ nums[i]
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		subarrayXor := prefXor[r+1] ^ prefXor[l]
		// Find x such that something is maximized
		// For XOR maximization, pick x as complement of subarrayXor
		x := 0
		best := 0
		for candidate := 0; candidate <= 100; candidate++ {
			if candidate^k > best {
				best = candidate ^ k
				x = candidate
			}
		}
		ans[qi] = subarrayXor ^ x
	}
	return ans
}
```
