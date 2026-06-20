# 1040 — Moving Stones Until Consecutive Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func numMovesStonesII(stones []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1040: Moving Stones Until Consecutive II
// https://leetcode.com/problems/moving-stones-until-consecutive-ii/
// Difficulty: Medium
//
// Approach: Sort stones. Max moves: spread left or right. Min moves: sliding window.
// Time: O(n log n)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStonesII([]int{7, 4, 9}))    // [1,2]
	fmt.Println(numMovesStonesII([]int{6, 5, 4, 3, 10})) // [2,3]
}

func numMovesStonesII(stones []int) []int {
  // Sort O(n log n)
	sort.Ints(stones)
	n := len(stones)

	// Max moves: fill gaps from one end, leaving one stone at the other end
	maxMoves := max(stones[n-1]-stones[1]-(n-2), stones[n-2]-stones[0]-(n-2))

	// Min moves: sliding window of size n
	minMoves := n
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i] < n {
			j++
		}
		already := j - i + 1
		moves := n - already
		// Special case: n-1 stones already consecutive, last one far away
		if moves == 1 && stones[j]-stones[i]+1 == n-1 && (j-i+1 == n-1) {
			moves = 2
		}
		if moves < minMoves {
			minMoves = moves
		}
	}

	return []int{minMoves, maxMoves}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
