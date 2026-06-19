package main

// LeetCode #2999: Count the Number of Powerful Integers
// https://leetcode.com/problems/count-the-number-of-powerful-integers/
// Difficulty: Hard
//
// Approach: Digit DP
// Count numbers <= X that end with suffix s and have all digits <= limit.
// countLe(X) returns how many powerful integers are in [1, X].
// Answer = countLe(finish) - countLe(start-1).

import (
	"fmt"
	"strconv"
)

func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	suffixVal, _ := strconv.ParseInt(s, 10, 64)
	countLe := func(x int64) int64 {
		if x < suffixVal {
			return 0
		}
		sx := strconv.FormatInt(x, 10)
		n := len(sx)
		m := len(s)
		if n < m {
			return 0
		}
		if n == m {
			if sx >= s {
				return 1
			}
			return 0
		}
		preLen := n - m
		var dfs func(pos int, tight bool) int64
		memo := make([][]int64, preLen)
		for i := range memo {
			memo[i] = []int64{-1, -1}
		}
		dfs = func(pos int, tight bool) int64 {
			if pos == preLen {
				if !tight {
					return 1
				}
				if sx[preLen:] >= s {
					return 1
				}
				return 0
			}
			ti := 0
			if tight {
				ti = 1
			}
			if memo[pos][ti] != -1 {
				return memo[pos][ti]
			}
			var res int64
			up := 9
			if tight {
				up = int(sx[pos] - '0')
			}
			if up > limit {
				up = limit
			}
			for d := 0; d <= up; d++ {
				res += dfs(pos+1, tight && d == int(sx[pos]-'0'))
			}
			memo[pos][ti] = res
			return res
		}
		ans := dfs(0, true)
		// Count numbers with fewer than n digits.
		// Numbers with exactly m digits (length of suffix) are handled by
		// leading zeros in dfs; so we start from m+1 digits up to n-1 digits.
		for length := m + 1; length < n; length++ {
			preLen2 := length - m
			if preLen2 == 0 {
				continue
			}
			ways := int64(limit)
			if limit > 9 {
				ways = 9
			}
			if limit >= 1 {
				for i := 1; i < preLen2; i++ {
					ways *= int64(limit + 1)
				}
				ans += ways
			}
		}
		return ans
	}
	return countLe(finish) - countLe(start-1)
}

func main() {
	// Example: [1,6000], limit=4, suffix="124" => 5
	fmt.Println(numberOfPowerfulInt(1, 6000, 4, "124"))
	// Example: [15,215], limit=6, suffix="10" => 3
	fmt.Println(numberOfPowerfulInt(15, 215, 6, "10"))
	// Single digit range
	fmt.Println(numberOfPowerfulInt(1, 10, 9, "5"))
}
