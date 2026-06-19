package main

// LeetCode #1808: Maximize Number of Nice Divisors
// https://leetcode.com/problems/maximize-number-of-nice-divisors/
// Difficulty: Hard
//
// Approach: Split primeFactors into groups of 3's for maximum product.
//   If n = primeFactors, we want to maximize the product of divisors,
//   which is equivalent to breaking n into positive integers that
//   multiply to the maximum value. Optimal strategy: use as many 3's
//   as possible. Handle remainders:
//     n % 3 == 0 -> 3^(n/3)
//     n % 3 == 1 -> 3^(n/3-1) * 4  (because 3*1 < 2*2)
//     n % 3 == 2 -> 3^(n/3) * 2
//   Result modulo 1_000_000_007.

import "fmt"

const MOD1808 = 1_000_000_007

func main() {
	// Example 1
	fmt.Println("Example 1 (p=5):", maxNiceDivisors(5))
	// Expected: 6  (n = 2*3 = 6, prime factors split into 2 and 3)

	// Example 2
	fmt.Println("Example 2 (p=8):", maxNiceDivisors(8))
	// Expected: 18 (split 8 = 3+3+2, product = 3*3*2 = 18)

	// Edge cases
	fmt.Println("Edge (p=1):", maxNiceDivisors(1))
	// Expected: 1
	fmt.Println("Edge (p=2):", maxNiceDivisors(2))
	// Expected: 2
	fmt.Println("Edge (p=3):", maxNiceDivisors(3))
	// Expected: 3
	fmt.Println("Edge (p=4):", maxNiceDivisors(4))
	// Expected: 4 (2*2)
}

func maxNiceDivisors(primeFactors int) int {
	if primeFactors <= 3 {
		return primeFactors
	}

	q := primeFactors / 3
	r := primeFactors % 3

	switch r {
	case 0:
		return modPow(3, q, MOD1808)
	case 1:
		// 3+1 -> 2+2 gives better product: 3*1 < 2*2
		return (modPow(3, q-1, MOD1808) * 4) % MOD1808
	default: // r == 2
		return (modPow(3, q, MOD1808) * 2) % MOD1808
	}
}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func MaximizeNumberOfNiceDivisors() any {
	return maxNiceDivisors(5)
}
