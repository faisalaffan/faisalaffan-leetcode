package main

// LeetCode #2827: Number of Beautiful Integers in the Range
// https://leetcode.com/problems/number-of-beautiful-integers-in-the-range/
// Difficulty: Hard
//
// Count integers in [low, high] where:
//   1) number of even digits == number of odd digits
//   2) the number is divisible by k
// Digit DP with memoization. State: (pos, mod, diff, started, tight).
// diff = odd_count - even_count (offset by +10 for 0-based indexing).
// O(len * k * 21 * 2 * 2 * 10) time, O(len * k * 21) space.

import (
	"fmt"
	"strconv"
)

func numberOfBeautifulIntegers(low int, high int, k int) int {
	return countLE(strconv.Itoa(high), k) - countLE(strconv.Itoa(low-1), k)
}

func countLE(s string, k int) int {
	n := len(s)
	digits := make([]int, n)
	for i, c := range s {
		digits[i] = int(c - '0')
	}

	memo := make([][][][]int, n)
	for i := range memo {
		memo[i] = make([][][]int, k)
		for j := range memo[i] {
			memo[i][j] = make([][]int, 21) // diff from -10 to +10, offset +10
			for l := range memo[i][j] {
				memo[i][j][l] = []int{-1, -1}
			}
		}
	}

	var dfs func(pos, mod, diff, started, tight int) int
	dfs = func(pos, mod, diff, started, tight int) int {
		if pos == n {
			if started == 1 && mod == 0 && diff == 10 {
				return 1
			}
			return 0
		}
		if memo[pos][mod][diff][tight] != -1 {
			return memo[pos][mod][diff][tight]
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
				// Leading zero: don't count digits, don't update mod
				total += dfs(pos+1, 0, 10, 0, ntight)
			} else {
				ndiff := diff
				if d%2 == 1 {
					ndiff++
				} else {
					ndiff--
				}
				total += dfs(pos+1, (mod*10+d)%k, ndiff, 1, ntight)
			}
		}

		memo[pos][mod][diff][tight] = total
		return total
	}

	return dfs(0, 0, 10, 0, 1)
}

func main() {
	// Example 1: [10,20], k=3 => 2 (12, 18)
	fmt.Println(numberOfBeautifulIntegers(10, 20, 3))

	// Example 2: [1,10], k=1 => 1 (10)
	fmt.Println(numberOfBeautifulIntegers(1, 10, 1))

	// Example 3: [5,5], k=2 => 0
	fmt.Println(numberOfBeautifulIntegers(5, 5, 2))

	// Single range
	fmt.Println(numberOfBeautifulIntegers(10, 10, 3)) // 10 -> odd=1,even=1, 10%3=1 -> 0
	fmt.Println(numberOfBeautifulIntegers(12, 12, 3)) // 12 -> odd=1,even=1, 12%3=0 -> 1

	// Small range
	fmt.Println(numberOfBeautifulIntegers(1, 100, 2))

	// k=1 (every number divisible by 1) - just count numbers with equal even/odd digits
	fmt.Println(numberOfBeautifulIntegers(10, 99, 1))

	// Single digit ranges (no beautiful numbers, can't have equal even/odd)
	fmt.Println(numberOfBeautifulIntegers(1, 9, 1))
}
