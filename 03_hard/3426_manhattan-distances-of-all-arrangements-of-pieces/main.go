package main

// LeetCode #3426: Manhattan Distances of All Arrangements of Pieces
// https://leetcode.com/problems/manhattan-distances-of-all-arrangements-of-pieces/
// Difficulty: Hard
//
// Math: contribution per cell pair = distance * C(m*n-2, k-2).
// Row and column contributions separate. MOD = 1e9+7.

import "fmt"

func main() {
	fmt.Println(ManhattanDistancesOfAllArrangementsOfPieces(2, 2, 2))
}

const MOD3426 = 1000000007

func powMod3426(a int64, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD3426
		}
		a = a * a % MOD3426
		b >>= 1
	}
	return res
}

func ManhattanDistancesOfAllArrangementsOfPieces(m, n, k int) int {
	total := int64(m) * int64(n)

	// Precompute factorials up to total
	fact := make([]int64, total+1)
	fact[0] = 1
	for i := int64(1); i <= total; i++ {
		fact[i] = fact[i-1] * i % MOD3426
	}
	invFact := make([]int64, total+1)
	invFact[total] = powMod3426(fact[total], MOD3426-2)
	for i := total - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD3426
	}

	nCr := func(nn, kk int64) int64 {
		if kk < 0 || kk > nn {
			return 0
		}
		return fact[nn] * invFact[kk] % MOD3426 * invFact[nn-kk] % MOD3426
	}

	rowContrib := int64(0)
	// Sum of distances between all pairs in same row of m rows, n cols
	// For each row: contribution = n*(n-1)*(n+1)/6 * m * C(total-2, k-2)
	rowPairs := int64(n) * int64(n-1) % MOD3426 * int64(n+1) % MOD3426
	rowPairs = rowPairs * powMod3426(6, MOD3426-2) % MOD3426
	rowContrib = rowPairs * int64(m) % MOD3426
	rowContrib = rowContrib * nCr(total-2, int64(k-2)) % MOD3426

	colContrib := int64(0)
	colPairs := int64(m) * int64(m-1) % MOD3426 * int64(m+1) % MOD3426
	colPairs = colPairs * powMod3426(6, MOD3426-2) % MOD3426
	colContrib = colPairs * int64(n) % MOD3426
	colContrib = colContrib * nCr(total-2, int64(k-2)) % MOD3426

	ans := (rowContrib + colContrib) % MOD3426
	return int(ans)
}
