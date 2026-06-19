package main

// LeetCode #3343: Count Number of Balanced Permutations
// https://leetcode.com/problems/count-number-of-balanced-permutations/
// Difficulty: Hard
//
// Count permutations where sum of digits at even indices = sum at odd indices.
// DP knapsack: for each digit d with frequency f, decide how many go to odd positions.
// Use combinatorics (factor + inverse factorial) for arranging each group.

import (
	"fmt"
)

func main() {
	// Example: n=2 -> 2 (balanced permutations of "01"... but n is just length)
	// Actually the problem uses a string of digits. Let me use the LeetCode format.
	// "12" -> 2 (permutations: "12" sum even=1, odd=2 not balanced; "21" sum even=2, odd=1)
	// Wait, the problem says n=2 -> 2
	fmt.Println(countBalancedPermutations("12"))

	// "123" -> 2
	fmt.Println(countBalancedPermutations("123"))

	// "112" -> 1
	fmt.Println(countBalancedPermutations("112"))

	// "12345" -> 0 (odd total sum)
	fmt.Println(countBalancedPermutations("12345"))

	// "0" -> 1
	fmt.Println(countBalancedPermutations("0"))
}

const MOD = 1000000007

func countBalancedPermutations(num string) int {
	n := len(num)
	cnt := make([]int, 10)
	totalSum := 0
	for _, ch := range num {
		d := int(ch - '0')
		cnt[d]++
		totalSum += d
	}

	if totalSum%2 != 0 {
		return 0
	}
	target := totalSum / 2
	oddPos := (n + 1) / 2 // ceil(n/2)

	// Precompute factorials and inverse factorials
	fact := make([]int, n+1)
	invFact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	invFact[n] = powMod(fact[n], MOD-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD
	}

	C := func(a, b int) int {
		if a < b || b < 0 {
			return 0
		}
		return fact[a] * invFact[b] % MOD * invFact[a-b] % MOD
	}

	// DP[k][s] = ways to choose k elements for odd positions summing to s
	dp := make([][]int, oddPos+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1

	for d := 0; d <= 9; d++ {
		f := cnt[d]
		if f == 0 {
			continue
		}
		// For each digit, try placing t copies in odd positions
		for k := oddPos; k >= 0; k-- {
			for s := target; s >= 0; s-- {
				if dp[k][s] == 0 {
					continue
				}
				for t := 1; t <= f && k+t <= oddPos && s+d*t <= target; t++ {
					ways := C(f, t)
					dp[k+t][s+d*t] = (dp[k+t][s+d*t] + dp[k][s]*ways) % MOD
				}
			}
		}
	}

	w := dp[oddPos][target]

	// ans = w * oddPos! * (n-oddPos)! / (prod cnt[d]!)
	ans := w * fact[oddPos] % MOD * fact[n-oddPos] % MOD
	for d := 0; d <= 9; d++ {
		ans = ans * invFact[cnt[d]] % MOD
	}
	return ans
}

func powMod(a, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		e >>= 1
	}
	return res
}
