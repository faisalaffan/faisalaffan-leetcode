package main

import (
	"fmt"
)

// 2318. Number of Distinct Roll Sequences
// ----------------------------------------------------------------
// Roll an n‑sided die (faces 1‑6) n times.  A sequence is valid if:
//   1. Adjacent rolls are different.
//   2. gcd(roll[i], roll[i-1]) == 1 for i ≥ 1.
//
// Count distinct sequences modulo 1e9+7.
//
// DP[i][prev] = number of valid sequences of length i ending with value prev.
// dp[i][p] = sum over q != p, gcd(p,q)==1 of dp[i-1][q].
//
// Because the transition depends only on the previous value we can
// pre‑compute which pairs are compatible, then run the DP.

const MOD = 1_000_000_007

func distinctRollSequences(n int) int {
	if n == 1 {
		return 6
	}

	// compatible[p][q] == true if q can follow p.
	compat := [7][7]bool{}
	for p := 1; p <= 6; p++ {
		for q := 1; q <= 6; q++ {
			if p != q && gcd(p, q) == 1 {
				compat[p][q] = true
			}
		}
	}

	dp := [7]int{}
	for p := 1; p <= 6; p++ {
		dp[p] = 1 // length 1
	}

	for i := 2; i <= n; i++ {
		ndp := [7]int{}
		for p := 1; p <= 6; p++ {
			s := 0
			for q := 1; q <= 6; q++ {
				if compat[q][p] {
					s = (s + dp[q]) % MOD
				}
			}
			ndp[p] = s
		}
		dp = ndp
	}

	sum := 0
	for p := 1; p <= 6; p++ {
		sum = (sum + dp[p]) % MOD
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
