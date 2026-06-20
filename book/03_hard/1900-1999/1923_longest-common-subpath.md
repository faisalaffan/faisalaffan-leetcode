# 1923 — Longest Common Subpath

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func longestCommonSubpath(n int, paths [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Binary Search

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1923: Longest Common Subpath
// https://leetcode.com/problems/longest-common-subpath/
// Difficulty: Hard
//
// Given n (number of cities) and paths (each is an array of cities visited),
// find the length of the longest subpath that appears in every path.
//
// Approach: binary search on length + rolling hash (Rabin-Karp).
// For a candidate length L, compute all hashes of subarrays of length L
// in the first path, then check if they appear in all other paths.
// Uses double hash (two moduli) to avoid collisions.

import (
	"fmt"
)

func main() {
	// Example 1:
	// n = 5, paths = [[0,1,2,3,4], [2,3,4], [4,0,1,2,3]]
	// Expected: 2 (subpath [2,3] or [3,4] or [4,0]... actually [2,3,4] doesn't appear in all)
	// [2,3] appears in all -> length 2
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}}))

	// Example 2:
	// n = 5, paths = [[0,1,2,3,4], [4,1,2,3,0], [1,2,3,4,0]]
	// Expected: 3 (subpath [1,2,3])
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1, 2, 3, 4}, {4, 1, 2, 3, 0}, {1, 2, 3, 4, 0}}))

	// Example 3:
	// n = 5, paths = [[0,0,0], [0,0,0]]
	// Expected: 3
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 0, 0}, {0, 0, 0}}))

	// Single path
	fmt.Println(longestCommonSubpath(3, [][]int{{1, 2, 3}}))

	// No common subpath
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1}, {2, 3}}))
}

func longestCommonSubpath(n int, paths [][]int) int {
	if len(paths) == 0 {
		return 0
	}
	if len(paths) == 1 {
		return len(paths[0])
	}

	// Find min path length for binary search upper bound
	low, high := 0, len(paths[0])
	for _, p := range paths {
		if len(p) < high {
			high = len(p)
		}
	}

	ans := 0
	for low <= high {
		mid := (low + high) / 2
		if hasCommonSubpath(paths, mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func hasCommonSubpath(paths [][]int, length int) bool {
	if length == 0 {
		return true
	}

	const mod1, mod2 = 1000000007, 1000000009
	const base = 100003

	// Precompute powers
	pow1, pow2 := 1, 1
	for i := 0; i < length-1; i++ {
		pow1 = pow1 * base % mod1
		pow2 = pow2 * base % mod2
	}

	// Get all subarray hashes from first path
  // HashMap: O(1) lookup
	common := make(map[[2]int]bool)
	var h1, h2 int

	for i, v := range paths[0] {
		x := v + 1 // shift by 1 to avoid hash issues with 0
		h1 = (h1*base + x) % mod1
		h2 = (h2*base + x) % mod2

		if i >= length {
			y := paths[0][i-length] + 1
			h1 = (h1 - y*pow1%mod1 + mod1) % mod1
			h2 = (h2 - y*pow2%mod2 + mod2) % mod2
		}

		if i >= length-1 {
			common[[2]int{h1, h2}] = true
		}
	}

	if len(common) == 0 {
		return false
	}

	// Check each remaining path
	for _, path := range paths[1:] {
  // HashMap: O(1) lookup
		current := make(map[[2]int]bool)
		h1, h2 = 0, 0

		for i, v := range path {
			x := v + 1
			h1 = (h1*base + x) % mod1
			h2 = (h2*base + x) % mod2

			if i >= length {
				y := path[i-length] + 1
				h1 = (h1 - y*pow1%mod1 + mod1) % mod1
				h2 = (h2 - y*pow2%mod2 + mod2) % mod2
			}

			if i >= length-1 {
				key := [2]int{h1, h2}
				if common[key] {
					current[key] = true
				}
			}
		}

		common = current
		if len(common) == 0 {
			return false
		}
	}
	return true
}
```
