# 3958 — Minimum Cost To Split Into Ones Ii

## Deskripsi

**Soal:** [3958. Minimum Cost To Split Into Ones Ii](https://leetcode.com/problems/minimum-cost-to-split-into-ones-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumCostToSplitIntoOnesIi(n int) int64`

> **Ide Kunci:** Optimal strategy splits off one 1 at a time.

## Solusi Go

```go
package main

// LeetCode #3958: Minimum Cost to Split into Ones II
// https://leetcode.com/problems/minimum-cost-to-split-into-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)
// Approach: Optimal strategy splits off one 1 at a time.
// Total cost = 1 + 2 + ... + (n-1) = n*(n-1)/2.

import "fmt"

func MinimumCostToSplitIntoOnesIi(n int) int64 {
	return int64(n) * int64(n-1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnesIi(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnesIi(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnesIi(1)) // Expected: 0
}
```
