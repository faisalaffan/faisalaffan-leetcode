package main

// LeetCode #1540: Can Convert String in K Moves
// https://leetcode.com/problems/can-convert-string-in-k-moves/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanConvertString("input", "ouput", 9))
	fmt.Println(CanConvertString("abc", "bcd", 10))
	fmt.Println(CanConvertString("aab", "bbb", 27))
}

func CanConvertString(s string, t string, k int) bool {
	// Time: O(N), Space: O(1)
	if len(s) != len(t) {
		return false
	}

	// Count how many times each shift value is needed
	shiftCount := make([]int, 26)

	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		// Compute needed shift (positive modulo 26)
		shift := (int(t[i]) - int(s[i]) + 26) % 26
		if shift == 0 {
			continue
		}

		// For each subsequent time we need this shift, add 26
		shiftCount[shift]++
		needed := shift + (shiftCount[shift]-1)*26
		if needed > k {
			return false
		}
	}

	return true
}
