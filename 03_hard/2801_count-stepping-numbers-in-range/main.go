package main

// LeetCode #2801: Count Stepping Numbers in Range
// https://leetcode.com/problems/count-stepping-numbers-in-range/
// Difficulty: Hard
//
// Digit DP. A stepping number has adjacent digits differing by exactly 1.
// Single-digit numbers (including 0) are stepping numbers.
// Count numbers in [low, high] inclusive. O(len * 10 * 2 * 2) time, O(len) space.

import "fmt"

const mod = 1000000007

func steppingNumbers(low, high string) int {
	if low == "0" {
		return countLE(high)
	}
	cHigh := countLE(high)
	cLow := countLE(decrement(low))
	return (cHigh - cLow + mod) % mod
}

func decrement(s string) string {
	b := []byte(s)
	i := len(b) - 1
	for i >= 0 && b[i] == '0' {
		b[i] = '9'
		i--
	}
	if i < 0 {
		return "0"
	}
	b[i]--
	if b[0] == '0' && len(b) > 1 {
		return string(b[1:])
	}
	return string(b)
}

// Count stepping numbers in [0, s] inclusive
func countLE(s string) int {
	n := len(s)
	digits := make([]int, n)
	for i, c := range s {
		digits[i] = int(c - '0')
	}

	dp := make([][][][]int, n+1)
	for i := range dp {
		dp[i] = make([][][]int, 10)
		for j := range dp[i] {
			dp[i][j] = make([][]int, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, 2)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	var dfs func(pos int, last int, tight int, started int) int
	dfs = func(pos int, last int, tight int, started int) int {
		if pos == n {
			return 1
		}
		if dp[pos][last][tight][started] != -1 {
			return dp[pos][last][tight][started]
		}

		limit := 9
		if tight == 1 {
			limit = digits[pos]
		}

		total := 0
		for d := 0; d <= limit; d++ {
			ntight := tight
			if tight == 1 && d < limit {
				ntight = 0
			}

			if started == 0 && d == 0 {
				total = (total + dfs(pos+1, 0, ntight, 0)) % mod
			} else if started == 0 {
				total = (total + dfs(pos+1, d, ntight, 1)) % mod
			} else {
				diff := d - last
				if diff < 0 {
					diff = -diff
				}
				if diff == 1 {
					total = (total + dfs(pos+1, d, ntight, 1)) % mod
				}
			}
		}

		dp[pos][last][tight][started] = total
		return total
	}

	return dfs(0, 0, 1, 0)
}

func main() {
	// Problem example: low="1", high="11" => 10
	fmt.Println(steppingNumbers("1", "11"))

	// Edge: 0 is a stepping number (single digit)
	fmt.Println(steppingNumbers("0", "0"))

	// All single digits: 0 through 9 = 10 numbers
	fmt.Println(steppingNumbers("0", "9"))

	// Two-digit range
	fmt.Println(steppingNumbers("10", "20"))

	// Same number
	fmt.Println(steppingNumbers("1", "1"))
	fmt.Println(steppingNumbers("5", "5"))

	// Larger range
	fmt.Println(steppingNumbers("90", "101"))

	// High=1 digit, low=1 digit (different)
	fmt.Println(steppingNumbers("3", "7"))

	// Large range test
	fmt.Println(steppingNumbers("0", "100"))

	// Stepping numbers around 100: 98, 101 are stepping (two-digit: all from
	// 10-98 stepping; three-digit: 101, 121, 123, 210, ...)
	fmt.Println(steppingNumbers("99", "123"))

	// Large numbers
	fmt.Println(steppingNumbers("1000", "2000"))

	// High number with many digits
	fmt.Println(steppingNumbers("0", "1000000"))
}
