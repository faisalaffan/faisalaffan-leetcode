package main

import (
	"fmt"
)

// 2338. Count the Number of Ideal Arrays
// ----------------------------------------------------------------
// An array of length n with values in [1, maxValue] is "ideal" if every
// element divides the next.
//
// Observation: the chain of *distinct* values in an ideal array forms a
// divisor chain.  Consecutive equal values are allowed while still satisfying
// divisibility (since v divides v).  Therefore an ideal array of length n
// corresponds to:
//   - Pick a divisor chain of length k (distinct values).
//   - Distribute n positions among the k distinct values, each appearing at
//     least once: C(n-1, k-1) ways (stars and bars).
//
// Let f(k) = number of divisor chains of length k with values in [1, maxValue].
//   f(1) = maxValue
//   f(k) = sum_{v=1}^{maxValue} dp[k][v]
//   where dp[k][v] = sum_{d|v, d<v} dp[k-1][d]
//
// dp[k][v] can be computed efficiently using harmonic series: for each d,
// for each multiple m of d (m = 2d, 3d, …, maxValue):
//   dp[k][m] += dp[k-1][d]
//
// The maximum chain length is at most about log₂(maxValue) (since each step
// at least doubles).  For maxValue ≤ 10⁴, max chain length ≤ 14.
//
// Answer = sum_{k=1}^{maxChain} C(n-1, k-1) * f(k) mod 1e9+7.

const MOD = 1_000_000_007

func idealArrays(n int, maxValue int) int {
	// Pre‑compute combinations up to C(n-1, 13) (max chain is ~14).
	maxK := 14
	if maxK > n {
		maxK = n
	}
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, maxK+1)
		C[i][0] = 1
		for j := 1; j <= i && j <= maxK; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}

	// dp[k][v] — we only need two layers at a time.
	dp := make([]int, maxValue+1)
	for v := 1; v <= maxValue; v++ {
		dp[v] = 1 // k = 1
	}

	f := make([]int, maxK+1)
	f[1] = maxValue

	for k := 2; k <= maxK; k++ {
		ndp := make([]int, maxValue+1)
		total := 0
		// For each d, add dp[d] to its multiples.
		for d := 1; d <= maxValue; d++ {
			if dp[d] == 0 {
				continue
			}
			for m := d * 2; m <= maxValue; m += d {
				ndp[m] = (ndp[m] + dp[d]) % MOD
			}
		}
		for v := 1; v <= maxValue; v++ {
			total = (total + ndp[v]) % MOD
		}
		f[k] = total
		dp = ndp
	}

	ans := 0
	for k := 1; k <= maxK; k++ {
		if f[k] == 0 {
			break
		}
		comb := C[n-1][k-1] // C(n-1, k-1)
		ans = (ans + comb*f[k]) % MOD
	}
	return ans
}

// ---------------------------------------------------------------------------
//  Wrapper

func CountTheNumberOfIdealArrays() interface{} {
	return idealArrays(5, 3)
}

func main() {
	fmt.Println(CountTheNumberOfIdealArrays())

	tests := []struct {
		n, maxV, want int
	}{
		{2, 5, 10},
		{5, 3, 11},
		{3, 2, 4}, // sequences of length 3 with values 1,2: 111,112,122,222 = 4
		{1, 10, 10},
		{4, 4, 20},
	}
	for _, tc := range tests {
		got := idealArrays(tc.n, tc.maxV)
		if got != tc.want {
			fmt.Printf("FAIL n=%d maxV=%d: got %d, want %d\n",
				tc.n, tc.maxV, got, tc.want)
		}
	}
	fmt.Println("Done testing 2338.")
}
