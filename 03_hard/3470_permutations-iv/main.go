package main

// LeetCode #3470: Permutations IV
// https://leetcode.com/problems/permutations-iv/
// Difficulty: Hard
//
// Count permutations of [1..n] with exactly k inversions AND no fixed points
// (i.e., derangements). Count mod 1_000_000_007.
// DP: dp[i][j] = count for permutations of size i with j inversions (derangements).

import "fmt"

const MOD = 1_000_000_007

func permutationsIV(n, k int) int {
	// dp[i][j] = count of derangements of size i with j inversions
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		// Prefix sums for sliding window
		prefix := make([]int, k+2)
		for j := 0; j <= k; j++ {
			prefix[j+1] = (prefix[j] + dp[i-1][j]) % MOD
		}
		for j := 0; j <= k; j++ {
			// When inserting element i into a permutation of size i-1:
			// Inserting at position p (0 = end, i-1 = front) adds p inversions.
			// But for derangements, element i cannot be at position i-1 (would be fixed).
			// Allowed insertions: positions 0..i-2 (i-1 positions)
			// Each adds 0..i-2 inversions respectively.
			// So dp[i][j] = sum_{p=0}^{i-2} dp[i-1][j-p] where p ≤ j
			// = sum of dp[i-1][j - (i-2) .. j]
			// Actually if j >= 0: sum of dp[i-1][max(0, j-(i-2)) .. j]
			left := j - (i - 2)
			if left < 0 {
				left = 0
			}
			right := j
			sum_ := prefix[right+1] - prefix[left]
			if sum_ < 0 {
				sum_ += MOD
			}
			dp[i][j] = sum_
		}
	}

	return dp[n][k]
}

func main() {
	// Test: n=3, k=1 -> expected 4
	fmt.Printf("n=3, k=1 -> %d (expected 4)\n", permutationsIV(3, 1))

	// Test: n=3, k=0 -> derangements of 3 with 0 inversions: only [3,1,2] which has 2 inversions.
	// [2,3,1] has 2 inversions. [3,2,1] has 3 inversions and isn't a derangement (2 fixed).
	// Derangements of 3: [2,3,1] (inv=2), [3,1,2] (inv=2). So k=0 -> 0.
	fmt.Printf("n=3, k=0 -> %d (expected 0)\n", permutationsIV(3, 0))

	// Test: n=3, k=2 -> derangements of 3 with 2 inversions: both have 2 inversions -> 2
	fmt.Printf("n=3, k=2 -> %d\n", permutationsIV(3, 2))

	// Test: n=4, k=3 -> ?
	fmt.Printf("n=4, k=3 -> %d\n", permutationsIV(4, 3))

	// Test: n=1, k=0 -> 0 (no derangement of size 1)
	fmt.Printf("n=1, k=0 -> %d (expected 0)\n", permutationsIV(1, 0))

	// Test: n=2, k=1 -> derangement of 2: [2,1] has 1 inversion -> 1
	fmt.Printf("n=2, k=1 -> %d (expected 1)\n", permutationsIV(2, 1))
}
