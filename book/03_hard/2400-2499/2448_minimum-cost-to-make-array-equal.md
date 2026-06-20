# 2448 — Minimum Cost To Make Array Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minCost(input [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2448: Minimum Cost to Make Array Equal
// https://leetcode.com/problems/minimum-cost-to-make-array-equal/
// Difficulty: Hard
//
// Weighted median. Consider each element nums[i] with weight cost[i].
// The optimal value to make all elements equal is the weighted median:
// find the point where cumulative cost crosses half of total cost.
// Time O(N log N) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minCost([][]int{{1, 3, 5, 2}, {2, 3, 1, 14}}))
	// Example 2
	fmt.Println(minCost([][]int{{2, 2, 2, 2, 2}, {4, 2, 8, 1, 3}}))
	// Single element
	fmt.Println(minCost([][]int{{1}, {1}}))
}

func minCost(input [][]int) int64 {
	nums, cost := input[0], input[1]
	n := len(nums)

	type pair struct {
		num, c int
	}
	pairs := make([]pair, n)
	totalCost := int64(0)
	for i := 0; i < n; i++ {
		pairs[i] = pair{nums[i], cost[i]}
		totalCost += int64(cost[i])
	}
  // Custom sort
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].num < pairs[j].num
	})

	// Find weighted median
	cum := int64(0)
	median := 0
	half := (totalCost + 1) / 2
	for _, p := range pairs {
		cum += int64(p.c)
		if cum >= half {
			median = p.num
			break
		}
	}

	// Compute cost to move all to median
	var ans int64
	for i := 0; i < n; i++ {
		diff := nums[i] - median
		if diff < 0 {
			diff = -diff
		}
		ans += int64(diff) * int64(cost[i])
	}
	return ans
}
```
