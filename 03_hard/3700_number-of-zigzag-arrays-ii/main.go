package main

// LeetCode #3700: Number of ZigZag Arrays II
// https://leetcode.com/problems/number-of-zigzag-arrays-ii/
// Difficulty: Hard
//
// Same as ZigZag Arrays I but n up to 1e9, l/r up to 75.
// Use matrix exponentiation.
//
// Approach: Build transition matrix of size 2*m x 2*m (m = r-l+1 <= 75)
// and use fast exponentiation.

import "fmt"

func main() {
	// Example 1
	fmt.Println(zigZagArraysII(3, 4, 5))
	// Example 2
	fmt.Println(zigZagArraysII(1000000000, 1, 3))
	// Edge: small n
	fmt.Println(zigZagArraysII(3, 1, 2))
}

func zigZagArraysII(n int, l int, r int) int {
	if n < 3 || l > r {
		return 0
	}
	m := r - l + 1
	if m < 2 {
		return 0
	}

	// DP with matrix exponentiation: state = 2*m values
	// First m: dp[val][0] (last move was increase, i.e., cur > prev)
	// Last m: dp[val][1] (last move was decrease, i.e., cur < prev)
	// Actually for n=1, each value has count 1 for both states
	// For each step, we transition:
	//   ndp[cur][0] = sum_{prev < cur} dp[prev][1]
	//   ndp[cur][1] = sum_{prev > cur} dp[prev][0]

	size := 2 * m
	mat := make([][]int64, size)
	for i := range mat {
		mat[i] = make([]int64, size)
	}

	for cur := 0; cur < m; cur++ {
		for prev := 0; prev < cur; prev++ {
			// From dp[prev][1] (decrease) to ndp[cur][0] (increase)
			mat[cur][prev+m] = 1 // prev+m is the index for dp[prev][1]
		}
		for prev := cur + 1; prev < m; prev++ {
			// From dp[prev][0] (increase) to ndp[cur][1] (decrease)
			mat[cur+m][prev] = 1 // prev is the index for dp[prev][0]
		}
	}

	// Initial vector for n=1
	init := make([]int64, size)
	for v := 0; v < m; v++ {
		init[v] = 1 // dp[v][0] = 1
		init[v+m] = 1 // dp[v][1] = 1
	}

	// Matrix power: mat^(n-1)
	power := matPow(mat, int64(n-1), MOD)
	// Multiply init * power
	result := mulVec(init, power, MOD)

	ans := int64(0)
	for i := 0; i < size; i++ {
		ans = (ans + result[i]) % MOD
	}
	return int(ans)
}

func matMul(A, B [][]int64, mod int64) [][]int64 {
	n := len(A)
	C := make([][]int64, n)
	for i := range C {
		C[i] = make([]int64, n)
		for k := 0; k < n; k++ {
			if A[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				C[i][j] = (C[i][j] + A[i][k]*B[k][j]) % mod
			}
		}
	}
	return C
}

func matPow(A [][]int64, power int64, mod int64) [][]int64 {
	n := len(A)
	R := make([][]int64, n)
	for i := range R {
		R[i] = make([]int64, n)
		R[i][i] = 1
	}
	for power > 0 {
		if power&1 == 1 {
			R = matMul(R, A, mod)
		}
		A = matMul(A, A, mod)
		power >>= 1
	}
	return R
}

func mulVec(v []int64, M [][]int64, mod int64) []int64 {
	n := len(v)
	r := make([]int64, n)
	for i := 0; i < n; i++ {
		var sum int64
		for k := 0; k < n; k++ {
			if v[k] == 0 {
				continue
			}
			sum = (sum + v[k]*M[k][i]) % mod
		}
		r[i] = sum
	}
	return r
}
