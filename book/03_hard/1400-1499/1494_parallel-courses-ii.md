# 1494 — Parallel Courses Ii

## Deskripsi

**Soal:** [1494. Parallel Courses Ii](https://leetcode.com/problems/parallel-courses-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

> **Ide Kunci:** DP over Bitmask

## Solusi Go

```go
package main

// LeetCode #1494: Parallel Courses II
// https://leetcode.com/problems/parallel-courses-ii/
// Difficulty: Hard
//
// Approach: DP over Bitmask
// prereq[mask] = bitmask of courses that are prerequisites for courses in mask.
// Actually, we precompute pre[c] = bitmask of direct prerequisites for course c.
// dp[mask] = minimum semesters to complete courses in mask.
// For each mask, compute available = all courses whose prerequisites are satisfied
// (i.e., pre[c] & mask == pre[c] for each c not in mask).
// Then try all subsets of available with size <= k, and transition:
// dp[mask | subset] = min(dp[mask | subset], dp[mask] + 1)

import "fmt"

func main() {
	// Example 1
	fmt.Println(minNumberOfSemesters(4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2))
	// Expected: 3

	// Example 2
	fmt.Println(minNumberOfSemesters(5, [][]int{{2, 1}, {3, 1}, {4, 1}, {5, 1}}, 2))
	// Expected: 3

	// Example 3
	fmt.Println(minNumberOfSemesters(11, [][]int{}, 2))
	// Expected: 6 (11 courses, 2 per semester = ceil(11/2) = 6)
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	// pre[c] = bitmask of direct prerequisites for course c (1-indexed)
  // Membuat slice untuk menyimpan hasil
	pre := make([]int, n)
	for _, dep := range dependencies {
		// dep[0] -> dep[1], meaning dep[1] has prerequisite dep[0]
		// We use 0-indexed internally
		pre[dep[1]-1] |= 1 << (dep[0] - 1)
	}

	total := 1 << n
	INF := n + 1
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	// Precompute required prerequisites for each mask (union of all pre[c] for c in mask)
  // Membuat slice untuk menyimpan hasil
	require := make([]int, total)
	for mask := 1; mask < total; mask++ {
		// find lowest set bit
		lsb := mask & -mask
		c := 0
		// find index of lsb
		for (1 << c) != lsb {
			c++
		}
		require[mask] = require[mask^lsb] | pre[c]
	}

	for mask := 0; mask < total; mask++ {
		if dp[mask] == INF {
			continue
		}
		// Courses that can be taken next: those whose prerequisites are satisfied
		// and that are not already taken.
		available := 0
		for c := 0; c < n; c++ {
			if mask&(1<<c) != 0 {
				continue
			}
			if pre[c]&^mask == 0 { // all prerequisites in mask
				available |= 1 << c
			}
		}

		if available == 0 {
			continue
		}

		// Try all subsets of available of size <= k
		sub := available
		for sub > 0 {
			if bitsCount(sub) <= k {
				next := mask | sub
				if dp[next] > dp[mask]+1 {
					dp[next] = dp[mask] + 1
				}
			}
			sub = (sub - 1) & available
		}
		// Also consider taking 0 courses? No, that's wasteful.
		// Actually the DP needs to consider taking available courses.
		// The loop above handles all non-empty subsets.
	}

	return dp[total-1]
}

func bitsCount(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x &= x - 1
	}
	return cnt
}
```
