# 0403 — Frog Jump

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canCross(stones []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #403: Frog Jump
// https://leetcode.com/problems/frog-jump/
// Difficulty: Hard
//
// A frog starts on stone 0 and jumps units of k. From stone i with a jump of
// size k, the next jump can be k-1, k, or k+1. Determine if the frog can
// reach the last stone.
// DP: map[stone]->set of jump sizes that can reach that stone.

import (
	"fmt"
)

func main() {
	// Example 1: [0,1,3,5,6,8,12,17] -> true
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	// Example 2: [0,1,2,3,4,8,9,11] -> false
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	// Edge: two stones
	fmt.Println(canCross([]int{0, 1}))
	// Edge: three stones, possible
	fmt.Println(canCross([]int{0, 1, 3}))
	// Edge: three stones, impossible (can't make jump 2 from stone 1)
	fmt.Println(canCross([]int{0, 2}))
}

func canCross(stones []int) bool {
	if len(stones) == 0 {
		return false
	}

	// Map stone position -> set of jump sizes
  // HashMap: O(1) lookup
	dp := make(map[int]map[int]bool, len(stones))
	for _, s := range stones {
		dp[s] = make(map[int]bool)
	}
	dp[stones[0]][0] = true // start with jump 0

	lastStone := stones[len(stones)-1]
  // HashMap: O(1) lookup
	stoneSet := make(map[int]bool, len(stones))
	for _, s := range stones {
		stoneSet[s] = true
	}

	for _, pos := range stones {
		for jump := range dp[pos] {
			for k := jump - 1; k <= jump+1; k++ {
				if k <= 0 {
					continue
				}
				next := pos + k
				if next == lastStone {
					return true
				}
				if stoneSet[next] {
					dp[next][k] = true
				}
			}
		}
	}

	return false
}
```
