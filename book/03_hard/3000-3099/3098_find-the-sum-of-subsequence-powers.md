# 3098 — Find The Sum Of Subsequence Powers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfPowers(nums []int, k int) int
```

> **💡 Hint:** DP with difference threshold

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n^2 * k + D * n * k)  
**Kompleksitas Ruang:** O(n * k)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3098: Find the Sum of Subsequence Powers
// https://leetcode.com/problems/find-the-sum-of-subsequence-powers/
// Difficulty: Hard
// Time: O(n^2 * k + D * n * k) | Space: O(n * k)
//
// Approach: DP with difference threshold
// 1. Sort nums. Collect distinct pairwise differences.
// 2. For each threshold d, count how many k-length subsequences have
//    minimum absolute difference >= d.
// 3. Answer = sum_{threshold d} countGe(d) * (d - prevDiff).
//    This uses the fact that countGe is piecewise-constant between
//    consecutive distinct differences.

import (
	"fmt"
	"sort"
)

const MOD = 1_000_000_007

func sumOfPowers(nums []int, k int) int {
	n := len(nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)

	// Collect distinct pairwise differences
  // Membuat map (HashMap) — pencarian O(1)
	diffSet := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diffSet[nums[j]-nums[i]] = true
		}
	}

  // Alokasi slice integer
	diffs := make([]int, 0, len(diffSet))
	for d := range diffSet {
		diffs = append(diffs, d)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(diffs)

	// countGe(threshold) = count of k-length subsequences whose min_abs_diff >= threshold
	countGe := func(threshold int) int {
  // Membuat matriks/slice 2D untuk DP
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, k+1)
		}

  // Alokasi slice integer
		runningSum := make([]int, k+1)
		ptr := 0

		for i := 0; i < n; i++ {
			// Two-pointer: advance ptr while difference from ptr to i >= threshold
			for ptr < i && nums[i]-nums[ptr] >= threshold {
				for j := 1; j <= k; j++ {
					runningSum[j] = (runningSum[j] + dp[ptr][j]) % MOD
				}
				ptr++
			}

			dp[i][1] = 1
			for j := 2; j <= k && j <= i+1; j++ {
				dp[i][j] = runningSum[j-1]
			}
		}

		total := 0
		for i := 0; i < n; i++ {
			total = (total + dp[i][k]) % MOD
		}
		return total
	}

	// Answer = sum_{d in diffs} countGe(d) * (d - prevDiff)
	answer := 0
	prevDiff := 0
	for _, d := range diffs {
		cnt := countGe(d)
		add := cnt * (d - prevDiff) % MOD
		answer = (answer + add) % MOD
		prevDiff = d
	}

	return answer
}

func main() {
	// Example 1
	fmt.Println("Test 1:", sumOfPowers([]int{1, 2, 3, 4}, 3))
	// Expected: 4
	// 3-length subsequences: [1,2,3] min diff=1, [1,2,4] min diff=1, [1,3,4] min diff=1, [2,3,4] min diff=1
	// Sum = 1+1+1+1 = 4

	// Example 2
	fmt.Println("Test 2:", sumOfPowers([]int{2, 2}, 2))
	// Expected: 0 (no 2-length subsequence with positive min diff since both elements are equal)

	// Example 3
	fmt.Println("Test 3:", sumOfPowers([]int{4, 3, -1}, 2))
	// Expected: 10
	// Sorted: [-1, 3, 4]
	// 2-length subsequences: [-1,3] diff=4, [-1,4] diff=5, [3,4] diff=1
	// Sum of powers = 4 + 5 + 1 = 10

	// k = 1
	fmt.Println("Test 4:", sumOfPowers([]int{1, 5, 10}, 1))
	// Expected: 0 (power of a single-element subsequence is undefined, typically 0)

	// Duplicates
	fmt.Println("Test 5:", sumOfPowers([]int{1, 1, 2}, 2))
	// Sorted: [1, 1, 2]
	// 2-length subsequences: [1,1] diff=0, [1,2] diff=1, [1,2] diff=1
	// Sum = 0 + 1 + 1 = 2

	// Larger example
	fmt.Println("Test 6:", sumOfPowers([]int{1, 3, 6, 10}, 2))
	// Sorted: [1, 3, 6, 10]
	// 2-length subsequences:
	// [1,3]=2, [1,6]=5, [1,10]=9, [3,6]=3, [3,10]=7, [6,10]=4
	// Sum = 2+5+9+3+7+4 = 30

	// Two elements, k=2
	fmt.Println("Test 7:", sumOfPowers([]int{5, 10}, 2))
	// Expected: 5 (diff = 5)

	// All same
	fmt.Println("Test 8:", sumOfPowers([]int{7, 7, 7}, 2))
	// Expected: 0 (all diffs are 0)
}
```
