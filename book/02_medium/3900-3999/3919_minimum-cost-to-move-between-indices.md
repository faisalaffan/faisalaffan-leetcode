# 3919 — Minimum Cost To Move Between Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func MinimumCostToMoveBetweenIndices(nums []int, queries [][]int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** O(N + Q)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3919: Minimum Cost to Move Between Indices
// https://leetcode.com/problems/minimum-cost-to-move-between-indices/
// Difficulty: Medium
// Time: O(N + Q) | Space: O(N)
// Approach: Build directed cost graph. From x to closest(x) costs 1, else
// abs(nums[x]-nums[y]). Since nums sorted, cheapest path is step-by-step
// using cost-1 edges when available. Precompute prefix sums for O(1) queries.

import "fmt"

func MinimumCostToMoveBetweenIndices(nums []int, queries [][]int) []int64 {
	n := len(nums)
	if n <= 1 {
  // Alokasi slice
		ans := make([]int64, len(queries))
		return ans
	}

	// costLR[i] = min cost from i to i+1
	// costRL[i] = min cost from i+1 to i
  // Alokasi slice
	costLR := make([]int64, n-1)
  // Alokasi slice
	costRL := make([]int64, n-1)

	for i := 0; i < n; i++ {
		if i == 0 {
			// closest(0) = 1
			costLR[0] = 1
		} else if i == n-1 {
			// closest(n-1) = n-2
			costRL[n-2] = 1
		} else {
			leftDiff := nums[i] - nums[i-1]
			rightDiff := nums[i+1] - nums[i]
			if leftDiff <= rightDiff {
				// closest(i) = i-1
				costRL[i-1] = 1
				costLR[i] = int64(rightDiff)
			} else {
				// closest(i) = i+1
				costLR[i] = 1
				costRL[i-1] = int64(leftDiff)
			}
		}
	}

	// prefLR[k] = cost from 0 to k (going right)
  // Alokasi slice
	prefLR := make([]int64, n)
	for i := 0; i < n-1; i++ {
		prefLR[i+1] = prefLR[i] + costLR[i]
	}

	// prefRL[k] = cost from k to 0 (going left)
  // Alokasi slice
	prefRL := make([]int64, n)
	for i := 0; i < n-1; i++ {
		prefRL[i+1] = prefRL[i] + costRL[i]
	}

  // Alokasi slice
	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l < r {
			ans[i] = prefLR[r] - prefLR[l]
		} else {
			ans[i] = prefRL[l] - prefRL[r]
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{-5, -2, 3},
		[][]int{{0, 2}, {2, 0}, {1, 2}},
	)) // Expected: [6 2 5]

	// Example 2
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{0, 2, 3, 9},
		[][]int{{3, 0}, {1, 2}, {2, 0}},
	)) // Expected: [4 1 3]

	// Example 3: single element (edge case)
	fmt.Println(MinimumCostToMoveBetweenIndices(
		[]int{5},
		[][]int{{0, 0}},
	)) // Expected: [0]
}
```
