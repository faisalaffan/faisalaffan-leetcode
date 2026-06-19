package main

// LeetCode #1259: Handshakes That Don't Cross
// https://leetcode.com/problems/handshakes-that-dont-cross/
// Difficulty: Hard [Paid]
//
// An even number of people sit in a circle. Each person shakes hands with
// exactly one other person, and no handshakes may cross. Count the number
// of ways this can happen.
//
// This is the (n/2)th Catalan number: C_k = (2k)! / (k! * (k+1)!)
// where k = n/2.

import "fmt"

func main() {
	// n = 2 (1 pair, 1 way): person 0 shakes with person 1
	fmt.Println(numberOfWays(2)) // 1

	// n = 4 (2 pairs, 2 ways): (0-1, 2-3) or (0-3, 1-2)
	// (0-2, 1-3) would cross
	fmt.Println(numberOfWays(4)) // 2

	// n = 6 (Catalan(3) = 5)
	fmt.Println(numberOfWays(6)) // 5

	// n = 8
	fmt.Println(numberOfWays(8)) // 14

	// n = 10
	fmt.Println(numberOfWays(10)) // 42

	// n = 12
	fmt.Println(numberOfWays(12)) // 132

	// n = 50 (large, tests overflow handling with mod)
	fmt.Println(numberOfWays(50))
}

const mod = 1000000007

// numberOfWays returns the number of non-crossing handshake configurations
// for n people (n is even), modulo 1_000_000_007.
//
// The solution uses DP with the Catalan recurrence derived from the first
// person shaking hands with person i (i must be odd for even spacing):
// dp[0] = 1
// dp[k] = sum(dp[i] * dp[k-1-i]) for i = 0..k-1
// where k = n/2.
func numberOfWays(n int) int {
	if n%2 != 0 {
		return 0
	}

	k := n / 2
	dp := make([]int, k+1)
	dp[0] = 1

	for i := 1; i <= k; i++ {
		total := 0
		for j := 0; j < i; j++ {
			total = (total + dp[j]*dp[i-1-j]) % mod
		}
		dp[i] = total
	}

	return dp[k]
}
