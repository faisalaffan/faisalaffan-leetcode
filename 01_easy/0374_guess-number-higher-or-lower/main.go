package main

// LeetCode #374: Guess Number Higher or Lower
// https://leetcode.com/problems/guess-number-higher-or-lower/
// Difficulty: Easy

import "fmt"

var pick int

func guess(num int) int {
	if num == pick {
		return 0
	} else if num < pick {
		return 1
	}
	return -1
}

// Time: O(log n), Space: O(1)
func GuessNumberHigherOrLower(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			lo = mid + 1
		case -1:
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	pick = 6
	fmt.Println(GuessNumberHigherOrLower(10))

	pick = 1
	fmt.Println(GuessNumberHigherOrLower(1))

	pick = 1
	fmt.Println(GuessNumberHigherOrLower(2))
}
