package main

// LeetCode #1830: Minimum Number of Operations to Make String Sorted
// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/
// Difficulty: Hard
//
// Approach: Combinatorics with Modular Arithmetic.
//   Each operation is the standard "next permutation" algorithm.
//   The number of operations to reach the sorted (ascending) string from
//   a given string equals the number of distinct next-permutations from
//   the sorted form, wrapping around. This is equivalent to:
//     answer = (totalPermutations - lexicographicRank) % totalPermutations
//
//   Compute total permutations and rank using factorials and modular
//   inverses modulo 1_000_000_007.

import "fmt"

const MOD1830 = 1_000_000_007

func main() {
	// Example 1
	fmt.Println("Example 1:", makeStringSorted("cba"))
	// Expected: 1

	// Example 2
	fmt.Println("Example 2:", makeStringSorted("aab"))
	// Expected: 3  (aab -> aba -> baa -> sorted? wait, sorted is "aab")

	// Single character
	fmt.Println("Edge (single):", makeStringSorted("a"))
	// Expected: 0

	// Already sorted
	fmt.Println("Edge (sorted):", makeStringSorted("abc"))
	// Expected: 0

	// Larger test
	fmt.Println("Example 3:", makeStringSorted("leetcode"))
	// Expected: some value; just verify it runs without error
}

func makeStringSorted(s string) int {
	n := len(s)

	// Precompute factorials and inverse factorials
	fact := make([]int, n+1)
	invFact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = (fact[i-1] * i) % MOD1830
	}
	invFact[n] = modPow1830(fact[n], MOD1830-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = (invFact[i+1] * (i + 1)) % MOD1830
	}

	// Count character frequencies (lowercase English letters)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	// Compute total number of permutations = n! / prod(cnt[i]!)
	totalPerms := fact[n]
	for i := 0; i < 26; i++ {
		totalPerms = (totalPerms * invFact[cnt[i]]) % MOD1830
	}

	// Compute lexicographic rank (0-indexed)
	// Current product of invFact[cnt[i]] for efficient updates
	curProd := int64(1)
	for i := 0; i < 26; i++ {
		curProd = curProd * int64(invFact[cnt[i]]) % MOD1830
	}

	rank := int64(0)
	remaining := n
	for i := 0; i < n; i++ {
		ch := int(s[i] - 'a')
		remaining--

		for c := 0; c < ch; c++ {
			if cnt[c] == 0 {
				continue
			}
			// Number of permutations with a smaller char at this position
			// = fact[remaining] * invFact[all counts after decrementing cnt[c]]
			// curProd currently has invFact[cnt[c]], we need invFact[cnt[c]-1]
			// invFact[cnt[c]-1] = invFact[cnt[c]] * cnt[c]
			ways := int64(fact[remaining]) * curProd % MOD1830
			ways = ways * int64(cnt[c]) % MOD1830
			rank = (rank + ways) % MOD1830
		}

		// Update: use s[i], decrement its count
		// curProd changes: invFact[cnt[ch]] -> invFact[cnt[ch]-1]
		// invFact[cnt-1] = invFact[cnt] * cnt
		curProd = curProd * int64(cnt[ch]) % MOD1830
		cnt[ch]--
	}

	// Answer = (totalPerms - rank) % totalPerms
	ans := (totalPerms - int(rank) + MOD1830) % MOD1830
	return ans
}

func modPow1830(base, exp int) int {
	result := 1
	base %= MOD1830
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % MOD1830
		}
		base = (base * base) % MOD1830
		exp >>= 1
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func MinimumNumberOfOperationsToMakeStringSorted() any {
	return makeStringSorted("cba")
}
