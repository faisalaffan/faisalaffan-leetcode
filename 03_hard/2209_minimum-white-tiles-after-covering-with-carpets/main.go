package main

// LeetCode #2209: Minimum White Tiles After Covering with Carpets
// https://leetcode.com/problems/minimum-white-tiles-after-covering-with-carpets/
// Difficulty: Hard
//
// DP: dp[i][j] = minimum white tiles for first i positions using j carpets.
// dp[i][j] = min(
//   dp[i-1][j] + cost(i),           // don't cover position i
//   dp[i-carpetLen][j-1]            // cover [i-carpetLen+1..i] with carpet
// )
// With 2-row optimization: prev = j-1, cur = j.

import (
	"fmt"
)

func main() {
	// "10110101", 2 carpets, carpetLen=2 => 2
	fmt.Println(minimumWhiteTiles("10110101", 2, 2))
	// All zeros
	fmt.Println(minimumWhiteTiles("00000", 2, 2))
	// All ones, 1 carpet of len 3
	fmt.Println(minimumWhiteTiles("11111", 1, 3))
	// Single tile
	fmt.Println(minimumWhiteTiles("1", 0, 1))
	fmt.Println(minimumWhiteTiles("0", 1, 1))
	// Extra: 1 carpet len 2
	fmt.Println(minimumWhiteTiles("10110101", 1, 2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	if n == 0 {
		return 0
	}

	// Row for 0 carpets: prefix count of white tiles.
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if floor[i-1] == '1' {
			dp[i]++
		}
	}

	for j := 1; j <= numCarpets; j++ {
		ndp := make([]int, n+1)
		// ndp[0] = 0 by default (0 positions, j carpets -> 0 white tiles)
		for i := 1; i <= n; i++ {
			// Option 1: don't place a carpet ending at i.
			// Use j carpets for first i-1 positions, add cost of position i.
			val := ndp[i-1]
			if floor[i-1] == '1' {
				val++
			}
			// Option 2: place a carpet covering [i-carpetLen+1, i].
			if i >= carpetLen {
				val = min(val, dp[i-carpetLen])
			} else {
				// Carpet extends before position 0 -- covers all of [0, i].
				val = min(val, 0)
			}
			ndp[i] = val
		}
		dp = ndp
	}

	return dp[n]
}
