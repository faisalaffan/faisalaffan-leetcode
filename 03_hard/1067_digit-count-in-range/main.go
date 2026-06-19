package main

// LeetCode #1067: Digit Count in Range
// https://leetcode.com/problems/digit-count-in-range/
// Difficulty: Hard [Paid]
//
// Count digit occurrences using mathematical decomposition per position
// (ones, tens, hundreds, ...). For a number n, countDigit(d, n) returns
// occurrences of digit d in [0, n]. Result for range [low, high] is
// countDigit(d, high) - countDigit(d, low-1).

import "fmt"

func main() {
	fmt.Println(digitCountInRange(1, 1, 13))
}

func digitCountInRange(d int, low int, high int) int {
	return countDigits(d, high) - countDigits(d, low-1)
}

func countDigits(d int, n int) int {
	if n < 0 {
		return 0
	}
	count := 0
	for pos := 1; pos <= n; pos *= 10 {
		left := n / (pos * 10)
		cur := (n / pos) % 10
		right := n % pos

		if d != 0 {
			if cur > d {
				count += (left + 1) * pos
			} else if cur == d {
				count += left*pos + right + 1
			} else {
				count += left * pos
			}
		} else {
			// digit 0: skip leading zeros
			if left > 0 {
				if cur > 0 {
					count += left * pos
				} else {
					count += (left-1)*pos + right + 1
				}
			}
		}
	}
	return count
}
