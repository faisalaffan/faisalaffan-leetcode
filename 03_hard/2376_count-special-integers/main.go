package main

// LeetCode #2376: Count Special Integers
// https://leetcode.com/problems/count-special-integers/
// Difficulty: Hard
//
// An integer is "special" if every digit in it is distinct.
// Count the number of special integers in the range [1, n].
//
// Approach: Digit DP with bitmask of used digits.
// dp[pos][mask][tight][started] = count of ways.
// pos: current position in the digit array
// mask: bitmask of digits already used
// tight: whether prefix matches n so far
// started: whether we've placed a non-leading-zero digit yet

import (
	"fmt"
	"strconv"
)

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)

	digits := make([]int, m)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	// memo[pos][mask][tight][started]
	memo := make([][][2][2]int, m)
	for i := range memo {
		memo[i] = make([][2][2]int, 1<<10)
		for mask := 0; mask < (1 << 10); mask++ {
			for t := 0; t < 2; t++ {
				for st := 0; st < 2; st++ {
					memo[i][mask][t][st] = -1
				}
			}
		}
	}

	var dp func(pos int, mask int, tight int, started int) int
	dp = func(pos int, mask int, tight int, started int) int {
		if pos == m {
			if started == 1 {
				return 1
			}
			return 0
		}

		if memo[pos][mask][tight][started] != -1 {
			return memo[pos][mask][tight][started]
		}

		limit := 9
		if tight == 1 {
			limit = digits[pos]
		}

		ans := 0
		for d := 0; d <= limit; d++ {
			nextStarted := started
			if d != 0 {
				nextStarted = 1
			}

			nextTight := tight
			if tight == 1 && d < limit {
				nextTight = 0
			}

			if nextStarted == 0 {
				// Still leading zeros, mask unchanged
				ans += dp(pos+1, mask, nextTight, nextStarted)
			} else {
				// Only use digit if not already used
				if mask&(1<<d) == 0 {
					ans += dp(pos+1, mask|(1<<d), nextTight, nextStarted)
				}
			}
		}

		memo[pos][mask][tight][started] = ans
		return ans
	}

	return dp(0, 0, 1, 0)
}

func main() {
	// Example 1
	fmt.Println(countSpecialNumbers(20))
	// Example 2
	fmt.Println(countSpecialNumbers(5))
	// Example 3
	fmt.Println(countSpecialNumbers(135))
	// Edge: 100
	fmt.Println(countSpecialNumbers(100))
	// Large number
	fmt.Println(countSpecialNumbers(1000))
	// Maximum constraint
	fmt.Println(countSpecialNumbers(987654321))
}
