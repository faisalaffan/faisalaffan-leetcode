# 3743 — Maximize Cyclic Partition Score

## Deskripsi

**Soal:** [3743. Maximize Cyclic Partition Score](https://leetcode.com/problems/maximize-cyclic-partition-score/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Duplicate array, run DP on 2n length. For each valid

## Solusi Go

```go
package main

// LeetCode #3743: Maximize Cyclic Partition Score
// https://leetcode.com/problems/maximize-cyclic-partition-score/
// Difficulty: Hard
//
// Partition a cyclic array into exactly k contiguous subarrays.
// Score = sum of (max - min) for each subarray. Maximize score.
//
// Approach: Duplicate array, run DP on 2n length. For each valid
// start position s (0 to n-1), compute linear DP on n elements
// partitioned into k subarrays. Use sparse table for O(1) range
// max/min queries.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(maximumScore([]int{2, 3, 3, 1}, 2))
	// Example 2
	fmt.Println(maximumScore([]int{1, 2, 3, 4}, 2))
	// Edge: k=1
	fmt.Println(maximumScore([]int{5, 1, 9, 3}, 1))
	// Edge: k=n
	fmt.Println(maximumScore([]int{1, 2, 3}, 3))
}

func maximumScore(nums []int, k int) int64 {
	n := len(nums)
	if n <= 0 || k <= 0 {
		return 0
	}
	if k == 1 {
		mx, mn := nums[0], nums[0]
		for _, v := range nums {
			if v > mx {
				mx = v
			}
			if v < mn {
				mn = v
			}
		}
		return int64(mx - mn)
	}
	if k >= n {
		return 0
	}

	// Build sparse table for range max and min queries
  // Membuat slice untuk menyimpan hasil
	log := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}
	K := log[n] + 1
  // Membuat slice 2D untuk DP/tabel
	stMax := make([][]int, K)
  // Membuat slice 2D untuk DP/tabel
	stMin := make([][]int, K)
  // Iterasi seluruh elemen
	for i := range stMax {
		stMax[i] = make([]int, n)
		stMin[i] = make([]int, n)
	}
	copy(stMax[0], nums)
	copy(stMin[0], nums)
	for j := 1; j < K; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			stMax[j][i] = max(stMax[j-1][i], stMax[j-1][i+(1<<(j-1))])
			stMin[j][i] = min(stMin[j-1][i], stMin[j-1][i+(1<<(j-1))])
		}
	}

	rangeMax := func(l, r int) int {
		j := log[r-l+1]
		return max(stMax[j][l], stMax[j][r-(1<<j)+1])
	}
	rangeMin := func(l, r int) int {
		j := log[r-l+1]
		return min(stMin[j][l], stMin[j][r-(1<<j)+1])
	}
	rangeScore := func(l, r int) int64 {
		return int64(rangeMax(l, r) - rangeMin(l, r))
	}

	// DP for linear array of size m partitioned into p segments
	// dp[m][p] = max score
	solveLinear := func(arr []int, m, p int) int64 {
		// dp[i][j] for first i elements (0..i-1), j segments
  // Membuat slice 2D untuk DP/tabel
		dp := make([][]int64, m+1)
  // Iterasi seluruh elemen
		for i := range dp {
			dp[i] = make([]int64, p+1)
			for j := range dp[i] {
				dp[i][j] = math.MinInt64
			}
		}
		dp[0][0] = 0

		for i := 1; i <= m; i++ {
			for j := 1; j <= p && j <= i; j++ {
				// Try all possible last segment starts [l, i-1]
				for l := j - 1; l < i; l++ {
					if dp[l][j-1] == math.MinInt64 {
						continue
					}
					score := dp[l][j-1] + rangeScore(l, i-1)
					if score > dp[i][j] {
						dp[i][j] = score
					}
				}
			}
		}
		return dp[m][p]
	}

	// Try each start position in the cycle
	var best int64 = math.MinInt64
	for start := 0; start < n; start++ {
  // Membuat slice untuk menyimpan hasil
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = nums[(start+i)%n]
		}
		score := solveLinear(arr, n, k)
		if score > best {
			best = score
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
