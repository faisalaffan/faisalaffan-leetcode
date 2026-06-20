package main

// LeetCode #3869: Count Fancy Numbers in a Range
// https://leetcode.com/problems/count-fancy-numbers-in-a-range/
// Difficulty: Hard
//
// Count numbers in [l, r] where adjacent digits alternate between
// odd and even parity.
//
// Approach: Digit DP. Track previous digit parity. Count numbers
// where each adjacent pair has different parity.

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1
	fmt.Println(countFancy(1, 100))
	// Example 2
	fmt.Println(countFancy(10, 20))
	// Edge: single digit
	fmt.Println(countFancy(5, 5))
	// Edge: large range
	fmt.Println(countFancy(100, 200))
}

func countFancy(l int64, r int64) int64 {
	if l > r {
		return 0
	}
	return countUpTo(r) - countUpTo(l-1)
}

func countUpTo(n int64) int64 {
	if n <= 0 {
		return 0
	}
	s := strconv.FormatInt(n, 10)
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	memo := make([][][][]int64, len(digits))
	for i := range memo {
		memo[i] = make([][][]int64, 2)
		for j := range memo[i] {
			memo[i][j] = make([][]int64, 2)
			for k := range memo[i][j] {
				memo[i][j][k] = make([]int64, 2)
				for p := range memo[i][j][k] {
					memo[i][j][k][p] = -1
				}
			}
		}
	}

	var dfs func(pos int, tight int, started int, prevParity int) int64
	dfs = func(pos int, tight int, started int, prevParity int) int64 {
		if pos == len(digits) {
			if started == 1 {
				return 1
			}
			return 0
		}
		if memo[pos][tight][started][prevParity] != -1 {
			return memo[pos][tight][started][prevParity]
		}

		limit := 9
		if tight == 1 {
			limit = digits[pos]
		}

		var total int64
		for d := 0; d <= limit; d++ {
			nt := 0
			if tight == 1 && d == limit {
				nt = 1
			}
			if started == 0 && d == 0 {
				total += dfs(pos+1, nt, 0, 0)
			} else {
				parity := d % 2
				if started == 0 {
					total += dfs(pos+1, nt, 1, parity)
				} else if parity != prevParity {
					total += dfs(pos+1, nt, 1, parity)
				}
			}
		}

		memo[pos][tight][started][prevParity] = total
		return total
	}

	return dfs(0, 1, 0, 0)
}
