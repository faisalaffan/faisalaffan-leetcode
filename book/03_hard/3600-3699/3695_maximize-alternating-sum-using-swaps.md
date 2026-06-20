# 3695 — Maximize Alternating Sum Using Swaps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxAlternatingSum(nums []int, swaps [][]int) int64
```

> **💡 Hint:** DSU to find connected components of swappable indices.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
	parent := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
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
  // Membuat map (HashMap) — pencarian O(1)
	compGroups := make(map[int][]int)
	for i := 0; i < n; i++ {
		r := find(i)
		compGroups[r] = append(compGroups[r], i)
	}

	result := int64(0)

	for _, indices := range compGroups {
		// Extract values
  // Alokasi slice integer
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
  // Alokasi slice integer
		evenVals := make([]int, len(evens))
  // Alokasi slice integer
		oddVals := make([]int, len(odds))
  // Range loop: iterasi dengan indeks + nilai
		for i := range evenVals {
			if i < len(vals) {
				evenVals[i] = vals[i]
			}
		}
  // Range loop: iterasi dengan indeks + nilai
		for i := range oddVals {
			if len(vals)-1-i >= 0 {
				oddVals[i] = vals[len(vals)-1-i]
			}
		}

		// Sort even indices and assign
  // Urutkan secara ascending — O(n log n)
		sort.Ints(evens)
		for i, idx := range evens {
			if i < len(evenVals) {
				nums[idx] = evenVals[i]
			}
		}
  // Urutkan secara ascending — O(n log n)
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
