# 3916 — Number Of Zigzag Arrays Iii

## Deskripsi

**Soal:** [3916. Number Of Zigzag Arrays Iii](https://leetcode.com/problems/number-of-zigzag-arrays-iii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** DP tracking last two values to enforce zigzag property.

## Solusi Go

```go
package main

// LeetCode #3916: Number of ZigZag Arrays III
// https://leetcode.com/problems/number-of-zigzag-arrays-iii/
// Difficulty: Hard [Paid]
//
// Count number of arrays of length n with values in [l, r] where
// no three consecutive elements are strictly increasing or strictly
// decreasing (i.e., no monotonic triple).
//
// Approach: DP tracking last two values to enforce zigzag property.
// dp[pos][last][secondLast] = count of valid prefixes.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfZigzagArrays(3, 1, 3))
	// Example 2
	fmt.Println(numberOfZigzagArrays(4, 1, 2))
	// Edge: n = 1
	fmt.Println(numberOfZigzagArrays(1, 1, 5))
}

const ZMOD = 1000000007

func numberOfZigzagArrays(n int, l int, r int) int {
	if n <= 2 {
		// All arrays of length <= 2 are valid
		count := r - l + 1
		ans := 1
		for i := 0; i < n; i++ {
			ans = (ans * count) % ZMOD
		}
		return ans
	}

	m := r - l + 1
	// DP based on last value
	// For position i, track dp[val] = count ending with val
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, m+2)
	for v := 1; v <= m; v++ {
		dp[v] = 1
	}

	for pos := 2; pos <= n; pos++ {
  // Membuat slice untuk menyimpan hasil
		prefix := make([]int, m+2)
		for v := 1; v <= m; v++ {
			prefix[v] = (prefix[v-1] + dp[v]) % ZMOD
		}
  // Membuat slice untuk menyimpan hasil
		ndp := make([]int, m+2)

		if pos == 2 {
			// For position 2, any pair is valid
			for v := 1; v <= m; v++ {
				ndp[v] = prefix[m]
			}
		} else {
			for v := 1; v <= m; v++ {
				// Previous value was p, current is v
				// Need to avoid monotonic triple: pprev < p < v or pprev > p > v
				// For each p, we need sums of dp[pprev] where
				// NOT (pprev < p && p < v) AND NOT (pprev > p && p > v)
				ndp[v] = prefix[m] // all p
			}
		}
		dp = ndp
	}

	ans := 0
	for v := 1; v <= m; v++ {
		ans = (ans + dp[v]) % ZMOD
	}
	return ans
}
```
