# 3695 — Maximize Alternating Sum Using Swaps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxAlternatingSum(nums []int, swaps [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting, Union-Find

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3695: Maximize Alternating Sum Using Swaps
// https://leetcode.com/problems/maximize-alternating-sum-using-swaps/
// Difficulty: Hard
//
// Given array nums and swap pairs (indices that can be swapped freely),
// maximize alternating sum: nums[0] - nums[1] + nums[2] - nums[3] + ...
//
// Approach: DSU to find connected components of swappable indices.
// Within each component, assign largest values to even indices (positive)
// and smallest values to odd indices (negative).

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(maxAlternatingSum([]int{1, 2, 3, 4}, [][]int{{0, 1}, {2, 3}}))
	// Example 2
	fmt.Println(maxAlternatingSum([]int{5, 3, 1, 7}, [][]int{{0, 2}, {1, 3}}))
	// Edge: no swaps
	fmt.Println(maxAlternatingSum([]int{1, 2, 3}, [][]int{}))
	// Edge: single element
	fmt.Println(maxAlternatingSum([]int{10}, [][]int{}))
}

func maxAlternatingSum(nums []int, swaps [][]int) int64 {
	n := len(nums)

	// DSU
  // Alokasi slice
	parent := make([]int, n)
  // Range loop
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	for _, s := range swaps {
		union(s[0], s[1])
	}

	// Group indices by component
  // HashMap: O(1) lookup
	compGroups := make(map[int][]int)
	for i := 0; i < n; i++ {
		r := find(i)
		compGroups[r] = append(compGroups[r], i)
	}

	result := int64(0)

	for _, indices := range compGroups {
		// Extract values
  // Alokasi slice
		vals := make([]int, len(indices))
		for i, idx := range indices {
			vals[i] = nums[idx]
		}
		sort.Sort(sort.Reverse(sort.IntSlice(vals)))

		// Separate even and odd indices
		var evens, odds []int
		for _, idx := range indices {
			if idx%2 == 0 {
				evens = append(evens, idx)
			} else {
				odds = append(odds, idx)
			}
		}

		// Assign largest values to even positions (positive contribution)
		// and smallest to odd positions (negative contribution)
  // Alokasi slice
		evenVals := make([]int, len(evens))
  // Alokasi slice
		oddVals := make([]int, len(odds))
  // Range loop
		for i := range evenVals {
			if i < len(vals) {
				evenVals[i] = vals[i]
			}
		}
  // Range loop
		for i := range oddVals {
			if len(vals)-1-i >= 0 {
				oddVals[i] = vals[len(vals)-1-i]
			}
		}

		// Sort even indices and assign
  // Sort O(n log n)
		sort.Ints(evens)
		for i, idx := range evens {
			if i < len(evenVals) {
				nums[idx] = evenVals[i]
			}
		}
  // Sort O(n log n)
		sort.Ints(odds)
		for i, idx := range odds {
			if i < len(oddVals) {
				nums[idx] = oddVals[i]
			}
		}
	}

	for i, v := range nums {
		if i%2 == 0 {
			result += int64(v)
		} else {
			result -= int64(v)
		}
	}
	return result
}
```
