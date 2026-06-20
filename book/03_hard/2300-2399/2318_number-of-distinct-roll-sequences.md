# 2318 — Number Of Distinct Roll Sequences

## Deskripsi

**Soal:** [2318. Number Of Distinct Roll Sequences](https://leetcode.com/problems/number-of-distinct-roll-sequences/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func distinctRollSequences(n int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// 2318. Number of Distinct Roll Sequences
// ----------------------------------------------------------------
// Roll a 6‑sided die n times.  A sequence is valid if:
//   1. gcd(roll[i], roll[i-1]) == 1 for i ≥ 1.
//   2. If roll[i] == roll[j] then |i - j| > 2  (gap of at least 2 between equals).
//
// Equivalent constraints used in DP:
//   - Adjacent: gcd(cur, prev) == 1  (implies cur != prev except for cur=prev=1).
//   - Gap-2:    cur != secPrev  (roll[i] != roll[i-2]).
//
// DP state: dp[prev][secPrev] = count of sequences ending with ..., secPrev, prev.
// Transition to cur: gcd(cur, prev) == 1 AND cur != secPrev.

const MOD = 1_000_000_007

func distinctRollSequences(n int) int {
	// Base: length 1 — any single value works.
	if n == 1 {
		return 6
	}

	// dp[prev][secPrev] — last two rolls of the current sequence.
	dp := [7][7]int{}

	// Initialise for length 2.
	for p := 1; p <= 6; p++ {
		for q := 1; q <= 6; q++ {
			if p != q && gcd(p, q) == 1 {
				dp[p][q] = 1 // sequence = [q, p]
			}
		}
	}

	if n == 2 {
		sum := 0
		for p := 1; p <= 6; p++ {
			for q := 1; q <= 6; q++ {
				sum = (sum + dp[p][q]) % MOD
			}
		}
		return sum
	}

	// Extend from length 3 to n.
	for i := 3; i <= n; i++ {
		ndp := [7][7]int{}
		for cur := 1; cur <= 6; cur++ {
			for prev := 1; prev <= 6; prev++ {
				if cur == prev || gcd(cur, prev) != 1 {
					continue
				}
				// Sum over all secPrev where (secPrev, prev) was valid at previous
				// step AND cur != secPrev (gap-2 constraint).
				s := 0
				for secPrev := 1; secPrev <= 6; secPrev++ {
					if cur == secPrev {
						continue
					}
					s = (s + dp[prev][secPrev]) % MOD
				}
				ndp[cur][prev] = s
			}
		}
		dp = ndp
	}

	sum := 0
	for cur := 1; cur <= 6; cur++ {
		for prev := 1; prev <= 6; prev++ {
			sum = (sum + dp[cur][prev]) % MOD
		}
	}
	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ---------------------------------------------------------------------------
//  Wrapper

func NumberOfDistinctRollSequences() interface{} {
	return distinctRollSequences(4)
}

func main() {
	fmt.Println(NumberOfDistinctRollSequences())

	// Tests from problem.
	cases := []struct{ n, want int }{
		{1, 6},
		{2, 22},
		{3, 66}, // from problem
		{4, 184},
	}
	for _, c := range cases {
		got := distinctRollSequences(c.n)
		if got != c.want {
			fmt.Printf("FAIL n=%d: got %d, want %d\n", c.n, got, c.want)
		}
	}
	fmt.Println("Done testing 2318.")
}
```
