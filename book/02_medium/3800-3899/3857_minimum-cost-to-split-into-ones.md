# 3857 — Minimum Cost To Split Into Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func MinimumCostToSplitIntoOnes(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3857: Minimum Cost to Split into Ones
// https://leetcode.com/problems/minimum-cost-to-split-into-ones/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Minimum cost = n*(n-1)/2. Equivalent to total edges in complete graph.

import "fmt"

func MinimumCostToSplitIntoOnes(n int) int {
	return n * (n - 1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnes(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnes(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnes(10))
}
```
