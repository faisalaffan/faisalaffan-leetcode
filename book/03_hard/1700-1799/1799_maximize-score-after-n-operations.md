# 1799 — Maximize Score After N Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxScore(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1799: Maximize Score After N Operations
// https://leetcode.com/problems/maximize-score-after-n-operations/
// Difficulty: Hard
//
// Approach: DP with Bitmask.
//   dp[mask] = max score from remaining elements represented by mask.
//   At each step, pick two unused elements, compute opNum * gcd(a,b),
//   and recurse on the new mask. Use memoization to avoid recomputation.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", maxScore([]int{1, 2, 3, 4, 5, 6}))
	// Expected: 14

	// Example 2
	fmt.Println("Example 2:", maxScore([]int{3, 4, 6, 8}))
	// Expected: 11

	// Edge case: 2 elements
	fmt.Println("Edge (2 elems):", maxScore([]int{1, 2}))
	// Expected: 1*1 = 1
}

func maxScore(nums []int) int {
	m := 1 << len(nums)
  // Alokasi slice
	memo := make([]int, m)
  // Range loop
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(mask int) int
	dfs = func(mask int) int {
		if memo[mask] != -1 {
			return memo[mask]
		}
		if mask == m-1 {
			return 0
		}
		bits := popcount(mask)
		op := bits/2 + 1
		best := 0
  // Linear scan O(n)
		for i := 0; i < len(nums); i++ {
			if mask&(1<<i) != 0 {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if mask&(1<<j) != 0 {
					continue
				}
				score := op*gcd(nums[i], nums[j]) + dfs(mask|(1<<i)|(1<<j))
				if score > best {
					best = score
				}
			}
		}
		memo[mask] = best
		return best
	}

	return dfs(0)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func popcount(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x &= x - 1
	}
	return cnt
}

// Stub kept for compatibility with the repo scaffold.
func MaximizeScoreAfterNOperations() any {
	return maxScore([]int{1, 2, 3, 4, 5, 6})
}
```
