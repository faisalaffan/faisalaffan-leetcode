package main

// LeetCode #3791: Number of Balanced Integers in a Range
// https://leetcode.com/problems/number-of-balanced-integers-in-a-range/
// Difficulty: Hard
//
// Count integers in [low, high] where sum of even-position digits
// equals sum of odd-position digits (positions from right, 1-indexed).
//
// Approach: Digit DP tracking positions and digit sum differences.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countBalanced(1, 100))
	// Example 2
	fmt.Println(countBalanced(10, 50))
	// Edge: single number
	fmt.Println(countBalanced(11, 11))
	// Edge: small range
	fmt.Println(countBalanced(1, 9))
}

func countBalanced(low int64, high int64) int64 {
	return countUpTo(high) - countUpTo(low-1)
}

func countUpTo(n int64) int64 {
	if n <= 0 {
		return 0
	}

	digits := split(n)
	m := len(digits)

	// dp[pos][diff][tight][started]
	// diff = sum_even - sum_odd (can be negative, offset by 9*m)
	offset := 9 * m
	memo := make([][][][]int64, m)
	for i := range memo {
		memo[i] = make([][][]int64, 2*offset+1)
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

	var dp func(pos int, diff int, tight int, started int) int64
	dp = func(pos int, diff int, tight int, started int) int64 {
		if pos == m {
			if started == 1 && diff == 0 {
				return 1
			}
			return 0
		}
		idx := diff + offset
		if idx < 0 || idx >= len(memo[pos]) {
			return 0
		}
		if memo[pos][idx][tight][started] != -1 {
			return memo[pos][idx][tight][started]
		}

		var total int64
		limit := 9
		if tight == 1 {
			limit = digits[pos]
		}

		for d := 0; d <= limit; d++ {
			nt := tight
			if tight == 1 && d < limit {
				nt = 0
			}
			if started == 0 && d == 0 {
				total += dp(pos+1, 0, nt, 0)
				continue
			}
			// Position from right: (m - 1 - pos) is the position from left
			// Position from right (1-indexed) = m - pos
			posFromRight := m - pos
			newDiff := diff
			if posFromRight%2 == 0 {
				// even position from right: add to sum_even
				newDiff += d
			} else {
				// odd position from right: subtract (sum_odd)
				newDiff -= d
			}
			total += dp(pos+1, newDiff, nt, 1)
		}

		memo[pos][idx][tight][started] = total
		return total
	}

	return dp(0, 0, 1, 0)
}

func split(n int64) []int {
	if n == 0 {
		return []int{0}
	}
	var d []int
	for n > 0 {
		d = append(d, int(n%10))
		n /= 10
	}
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	return d
}
