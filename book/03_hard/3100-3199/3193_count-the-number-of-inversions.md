# 3193 — Count The Number Of Inversions

## Deskripsi

**Soal:** [3193. Count The Number Of Inversions](https://leetcode.com/problems/count-the-number-of-inversions/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func numberOfPermutations(n int, requirements [][]int) int`

> **Ide Kunci:** DP with prefix sums.

## Solusi Go

```go
package main

// LeetCode #3193: Count the Number of Inversions
// https://leetcode.com/problems/count-the-number-of-inversions/
// Difficulty: Hard
//
// Count permutations of [0, n-1] satisfying inversion-count requirements.
// requirements: [(i, cnt)] means prefix ending at index i must have exactly
// cnt inversions.
//
// Approach: DP with prefix sums.
//   dp[i][j] = number of ways for first i+1 elements with j inversions.
//   dp[i][j] = sum(dp[i-1][j-k] for k=0..min(i, j)).
// Use prefix sums for O(n^2) time.

import (
	"fmt"
)

const MOD = 1000000007

func numberOfPermutations(n int, requirements [][]int) int {
	// required[i] = required inversion count for prefix ending at i, or -1 if
	// unspecified.
  // Membuat slice untuk menyimpan hasil
	required := make([]int, n)
  // Iterasi seluruh elemen
	for i := range required {
		required[i] = -1
	}
	maxInv := 0
	for _, req := range requirements {
		idx, cnt := req[0], req[1]
		required[idx] = cnt
		if cnt > maxInv {
			maxInv = cnt
		}
	}

	// dp[j] = number of ways to have exactly j inversions for current prefix.
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, maxInv+1)
	dp[0] = 1

	for i := 1; i < n; i++ {
		// prefix sums of dp for sliding window of size i+1 (0 to i).
  // Membuat slice untuk menyimpan hasil
		pref := make([]int, maxInv+2)
		pref[0] = dp[0]
		for j := 1; j <= maxInv; j++ {
			pref[j] = (pref[j-1] + dp[j]) % MOD
		}

  // Membuat slice untuk menyimpan hasil
		ndp := make([]int, maxInv+1)
		maxJ := maxInv
		if required[i] != -1 {
			maxJ = required[i]
		}

		for j := 0; j <= maxJ; j++ {
			// sum of dp[j-k] for k = 0..min(i, j)
			// = pref[j] - pref[j-min(i,j)-1]
			low := j - min(i, j) - 1
			val := pref[j]
			if low >= 0 {
				val = (val - pref[low] + MOD) % MOD
			}
			ndp[j] = val
		}

		dp = ndp
	}

	return dp[required[n-1]] % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numberOfPermutations(3, [][]int{{2, 2}})) // expect 1 (only [2,1,0])
	fmt.Println(numberOfPermutations(3, [][]int{{2, 0}})) // expect 1 (only [0,1,2])
}
```
