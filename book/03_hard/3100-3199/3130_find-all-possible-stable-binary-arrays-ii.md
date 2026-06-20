# 3130 — Find All Possible Stable Binary Arrays Ii

## Deskripsi

**Soal:** [3130. Find All Possible Stable Binary Arrays Ii](https://leetcode.com/problems/find-all-possible-stable-binary-arrays-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func numberOfStableArrays(zero, one, limit int) int`

## Solusi Go

```go
package main

// LeetCode #3130: Find All Possible Stable Binary Arrays II
// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-ii/
// Difficulty: Hard
//
// Count binary arrays of length zero+one with exactly zero zeros and one ones,
// such that no more than `limit` consecutive same elements appear.
// Uses DP with sliding window prefix sums for O(zero * one) time.
//
// dp0[i][j] = ways ending with 0, using i zeros and j ones
// dp1[i][j] = ways ending with 1, using i zeros and j ones
//
// dp0[i][j] = sum_{k=1}^{min(limit,i)} dp1[i-k][j]
// dp1[i][j] = sum_{k=1}^{min(limit,j)} dp0[i][j-k]

import (
	"fmt"
)

const MOD = 1000000007

func numberOfStableArrays(zero, one, limit int) int {
  // Membuat slice 2D untuk DP/tabel
	dp0 := make([][]int, zero+1)
  // Membuat slice 2D untuk DP/tabel
	dp1 := make([][]int, zero+1)
	// Prefix sums for sliding window optimization: pref0[i][j] = sum_{a<=i} dp0[a][j]
  // Membuat slice 2D untuk DP/tabel
	pref0 := make([][]int, zero+1)
	// We need row-wise prefix for dp1: pref1Row[i][j] = sum_{b<=j} dp1[i][b]
  // Membuat slice 2D untuk DP/tabel
	pref1Row := make([][]int, zero+1)

	for i := 0; i <= zero; i++ {
		dp0[i] = make([]int, one+1)
		dp1[i] = make([]int, one+1)
		pref0[i] = make([]int, one+1)
		pref1Row[i] = make([]int, one+1)
	}

	// Base cases: all zeros or all ones
	for k := 1; k <= limit && k <= zero; k++ {
		dp0[k][0] = 1
	}
	for k := 1; k <= limit && k <= one; k++ {
		dp1[0][k] = 1
	}

	// Build prefix sums
	for j := 0; j <= one; j++ {
		for i := 0; i <= zero; i++ {
			if i == 0 {
				pref0[i][j] = dp0[i][j]
			} else {
				pref0[i][j] = (pref0[i-1][j] + dp0[i][j]) % MOD
			}
		}
	}
	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			if j == 0 {
				pref1Row[i][j] = dp1[i][j]
			} else {
				pref1Row[i][j] = (pref1Row[i][j-1] + dp1[i][j]) % MOD
			}
		}
	}

	// Fill DP tables row by row (increasing total length)
	for total := 1; total <= zero+one; total++ {
		for i := 0; i <= zero && i <= total; i++ {
			j := total - i
			if j > one || j < 0 {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}

			if i > 0 {
				// dp0[i][j] = sum_{k=1}^{min(limit,i)} dp1[i-k][j]
				lo := i - limit
				if lo < 0 {
					lo = 0
				}
				val := pref1Row[i-1][j]
				if lo > 0 {
					val = (val - pref1Row[lo-1][j] + MOD) % MOD
				}
				dp0[i][j] = val
			}
			if j > 0 {
				// dp1[i][j] = sum_{k=1}^{min(limit,j)} dp0[i][j-k]
				lo := j - limit
				if lo < 0 {
					lo = 0
				}
				val := pref0[i][j-1]
				if lo > 0 {
					val = (val - pref0[i][lo-1] + MOD) % MOD
				}
				dp1[i][j] = val
			}

			// Update column prefix for dp0
			if i == 0 {
				pref0[i][j] = dp0[i][j]
			} else {
				pref0[i][j] = (pref0[i-1][j] + dp0[i][j]) % MOD
			}
			// Update row prefix for dp1
			if j == 0 {
				pref1Row[i][j] = dp1[i][j]
			} else {
				pref1Row[i][j] = (pref1Row[i][j-1] + dp1[i][j]) % MOD
			}
		}
	}

	return (dp0[zero][one] + dp1[zero][one]) % MOD
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfStableArrays(1, 1, 2))
	// Expected: 2 ([0,1] and [1,0])

	// Test case 2
	fmt.Println("Test 2:", numberOfStableArrays(1, 2, 1))
	// Expected: 1 ([1,0,1])

	// Test case 3
	fmt.Println("Test 3:", numberOfStableArrays(3, 1, 1))
	// Expected: 0 (can't place 3 zeros with limit=1)

	// Test case 4
	fmt.Println("Test 4:", numberOfStableArrays(2, 2, 1))
	// Expected: 2 ([0,1,0,1] and [1,0,1,0])

	// Test case 5: larger
	fmt.Println("Test 5:", numberOfStableArrays(3, 3, 2))
	// Expected: some number > 0

	// Test case 6: single element
	fmt.Println("Test 6:", numberOfStableArrays(1, 0, 5))
	// Expected: 1 ([0])
}
```
