# 2809 — Minimum Time To Make Array Sum At Most X

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumTime(nums1, nums2 []int, x int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2809: Minimum Time to Make Array Sum At Most x
// https://leetcode.com/problems/minimum-time-to-make-array-sum-at-most-x/
// Difficulty: Hard
//
// Each second: all nums1[i] += nums2[i], then you may zero out one element.
// Find minimum seconds to make sum(nums1) <= x, or -1 if impossible.
// Sort by growth rate (nums2), then DP[t] = max total reduction with t ops.
// O(N^2) time, O(N) space.

import (
	"fmt"
	"sort"
)

func minimumTime(nums1, nums2 []int, x int) int {
	n := len(nums1)
  // Alokasi slice
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}
  // Custom sort
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })

	// dp[t] = max total reduction achievable with exactly t resets
  // Alokasi slice
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		a, b := pairs[i][0], pairs[i][1]
		for t := i + 1; t >= 1; t-- {
			cand := dp[t-1] + a + b*t
			if cand > dp[t] {
				dp[t] = cand
			}
		}
	}

	sum1, sum2 := 0, 0
	for i := 0; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
	}

	for t := 0; t <= n; t++ {
		if sum1+sum2*t-dp[t] <= x {
			return t
		}
	}
	return -1
}

func main() {
	// Example: possible in 1 second
	fmt.Println(minimumTime([]int{1, 2, 3}, []int{1, 1, 1}, 4))

	// Need multiple seconds
	fmt.Println(minimumTime([]int{1, 2, 3}, []int{3, 3, 3}, 4))

	// Already <= x
	fmt.Println(minimumTime([]int{1, 2, 3}, []int{1, 1, 1}, 10))

	// Single element
	fmt.Println(minimumTime([]int{5}, []int{2}, 3))

	// Impossible (sum always > x)
	fmt.Println(minimumTime([]int{10, 10}, []int{100, 100}, 5))

	// Larger example
	fmt.Println(minimumTime([]int{3, 5, 7, 9}, []int{2, 2, 2, 2}, 15))

	// All same values
	fmt.Println(minimumTime([]int{10, 10, 10}, []int{1, 1, 1}, 25))

	// x is very large
	fmt.Println(minimumTime([]int{100, 200}, []int{10, 20}, 100000))
}
```
