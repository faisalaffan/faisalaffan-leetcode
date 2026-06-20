package main

// LeetCode #3704: Count No-Zero Pairs That Sum to N
// https://leetcode.com/problems/count-no-zero-pairs-that-sum-to-n/
// Difficulty: Hard
//
// Count pairs (a, b) such that a + b = n, a <= b, and both a and b have
// no zero digit in their decimal representation.
//
// Approach: Digit DP with carry/borrow tracking.

import "fmt"
import "strconv"

func main() {
	// Example 1
	fmt.Println(countNoZeroPairs(10))
	// Example 2
	fmt.Println(countNoZeroPairs(2))
	// Example 3
	fmt.Println(countNoZeroPairs(100))
	// Edge: single digit
	fmt.Println(countNoZeroPairs(5))
}

func countNoZeroPairs(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}
	m := len(digits)

	// memo[pos][carry][aStarted][aLess][bLess]
	// carry: 0 or 1 (from previous digit addition)
	// aStarted: whether a has started (non-zero digit seen)
	// aLess: whether a is already less than its bound
	// bLess: whether b is already less than its bound
	memo := make([][][][][]int64, m+1)
	for i := range memo {
		memo[i] = make([][][][]int64, 2)
		for j := range memo[i] {
			memo[i][j] = make([][][]int64, 2)
			for k := range memo[i][j] {
				memo[i][k] = make([][]int64, 2)
				for l := range memo[i][k] {
					memo[i][k][l] = make([]int64, 2)
					for p := range memo[i][k][l] {
						memo[i][k][l][p] = -1
					}
				}
			}
		}
	}

	var dp func(pos int, carry int, aStarted int, aLess int, bLess int) int64
	dp = func(pos int, carry int, aStarted int, aLess int, bLess int) int64 {
		if pos == m {
			if carry == 0 && aStarted == 1 {
				return 1
			}
			return 0
		}
		if memo[pos][carry][aStarted][aLess][bLess] != -1 {
			return memo[pos][carry][aStarted][aLess][bLess]
		}

		var total int64
		digit := digits[m-1-pos]

		for aD := 0; aD <= 9; aD++ {
			if !(aLess == 1) && aD > digit {
				break
			}
			if aStarted == 0 && aD == 0 {
				// a hasn't started, skip
				continue
			}
			for bD := 0; bD <= 9; bD++ {
				// b must be >= a for a <= b
				// but at this digit level, we can compare when a and b have same length
				if aD == 0 || bD == 0 {
					continue
				}
				sum := aD + bD + carry
				newCarry := sum / 10
				if sum%10 != digit {
					continue
				}
				newALess := aLess
				if aD < digit {
					newALess = 1
				}
				newBLess := bLess
				if bD < digit {
					newBLess = 1
				}
				total += dp(pos+1, newCarry, 1, newALess, newBLess)
			}
		}

		memo[pos][carry][aStarted][aLess][bLess] = total
		return total
	}

	return dp(0, 0, 0, 0, 0)
}
