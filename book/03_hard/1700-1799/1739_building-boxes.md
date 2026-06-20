# 1739 — Building Boxes

## Deskripsi

**Soal:** [1739. Building Boxes](https://leetcode.com/problems/building-boxes/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal), Stack (tumpukan LIFO)

> **Ide Kunci:** Greedy accumulation. Build a tetrahedral layer structure,

## Solusi Go

```go
package main

// LeetCode #1739: Building Boxes
// https://leetcode.com/problems/building-boxes/
// Difficulty: Hard
//
// You have n boxes to place on the floor. The boxes must be placed
// such that:
// - Each box is at the corner of a unit cube grid.
// - Boxes can be stacked, but each box above must be supported.
// Find the minimum number of boxes touching the floor.
//
// Approach: Greedy accumulation. Build a tetrahedral layer structure,
// counting boxes placed on the floor minimally.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumBoxes(3))
	// Example 2
	fmt.Println(minimumBoxes(4))
	// Example 3
	fmt.Println(minimumBoxes(10))
	// Edge: n = 1
	fmt.Println(minimumBoxes(1))
}

func minimumBoxes(n int) int {
	s, k := 0, 1
	for s+k*(k+1)/2 <= n {
		s += k * (k + 1) / 2
		k++
	}
	k--
	ans := k * (k + 1) / 2
	k = 1
	for s < n {
		ans++
		s += k
		k++
	}
	return ans
}
```
