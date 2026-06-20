# 1601 — Maximum Number Of Achievable Transfer Requests

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumRequests(n int, requests [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Bitmask** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1601: Maximum Number of Achievable Transfer Requests
// https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/
// Difficulty: Hard
//
// Bitmask enumeration approach:
// - There are at most 20 requests and 20 buildings (n <= 20).
// - Try all subsets from largest to smallest (or use bitmask enumeration).
// - For each subset, simulate the transfers and check if net change is 0.
// - Return the size of the largest valid subset.

import "fmt"

func main() {
	// Example: n=5, requests=[[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]] -> 5
	fmt.Println(maximumRequests(5, [][]int{
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4},
	}))

	// Additional tests
	fmt.Println(maximumRequests(3, [][]int{
		{0, 0},
		{1, 2},
		{2, 1},
	}))

	fmt.Println(maximumRequests(4, [][]int{
		{0, 1},
		{1, 0},
		{2, 3},
		{3, 2},
	}))

	fmt.Println(maximumRequests(3, [][]int{
		{0, 1},
		{1, 0},
		{1, 0},
		{0, 1},
	}))

	fmt.Println(maximumRequests(2, [][]int{
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 0},
	}))
}

func maximumRequests(n int, requests [][]int) int {
	r := len(requests)
	result := 0

	// Try all subsets
	for mask := 1; mask < (1 << r); mask++ {
		size := popcount(mask)
		if size <= result {
			continue
		}
		if isValid(mask, n, requests) {
			result = size
		}
	}

	return result
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func isValid(mask int, n int, requests [][]int) bool {
  // Alokasi slice
	balance := make([]int, n)
	for i, req := range requests {
		if mask&(1<<i) != 0 {
			balance[req[0]]--
			balance[req[1]]++
		}
	}
	for _, b := range balance {
		if b != 0 {
			return false
		}
	}
	return true
}
```
