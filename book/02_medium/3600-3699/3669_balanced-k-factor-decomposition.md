# 3669 — Balanced K Factor Decomposition

## Deskripsi

**Soal:** [3669. Balanced K Factor Decomposition](https://leetcode.com/problems/balanced-k-factor-decomposition/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(d^k) worst case with pruning  
**Kompleksitas Ruang:** O(k)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func balancedKFactorDecomposition(n int, k int) []int`

## Solusi Go

```go
package main

// LeetCode #3669: Balanced K-Factor Decomposition
// https://leetcode.com/problems/balanced-k-factor-decomposition/
// Difficulty: Medium
// Time: O(d^k) worst case with pruning | Space: O(k)

import (
	"fmt"
	"math"
)

func balancedKFactorDecomposition(n int, k int) []int {
	bestDiff := math.MaxInt32
	var best []int
  // Membuat slice untuk menyimpan hasil
	path := make([]int, k)

	var dfs func(rem int, start int, depth int)
	dfs = func(rem int, start int, depth int) {
		if depth == k-1 {
			if rem >= start {
				path[depth] = rem
				mn, mx := path[0], path[0]
				for _, v := range path {
					if v < mn {
						mn = v
					}
					if v > mx {
						mx = v
					}
				}
				diff := mx - mn
				if diff < bestDiff {
					bestDiff = diff
					best = make([]int, k)
					copy(best, path)
				}
			}
			return
		}

		for d := start; d*d <= rem; d++ {
			if rem%d == 0 {
				path[depth] = d
				dfs(rem/d, d, depth+1)
			}
		}
	}

	dfs(n, 1, 0)
	return best
}

func main() {
	fmt.Println(balancedKFactorDecomposition(100, 2))
	fmt.Println(balancedKFactorDecomposition(44, 3))
	fmt.Println(balancedKFactorDecomposition(12, 2))
}
```
