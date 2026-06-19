package main

// LeetCode #1542: Find Longest Awesome Substring
// https://leetcode.com/problems/find-longest-awesome-substring/
// Difficulty: Hard
//
// Bitmask prefix approach:
// - A substring is "awesome" if at most one digit has odd frequency.
// - Use a 10-bit mask where bit k = parity of digit k's count.
// - For each prefix position, store the first occurrence of each mask.
// - For each mask, we look for the same mask (all even) or
//   a mask differing by exactly one bit (one digit odd).

import "fmt"

func main() {
	// Example: "3242415" -> 5 ("24241" or "42415")
	fmt.Println(longestAwesome("3242415"))

	// Additional tests
	fmt.Println(longestAwesome("0"))
	fmt.Println(longestAwesome("00"))
	fmt.Println(longestAwesome("123456789"))
	fmt.Println(longestAwesome("373781"))
}

func longestAwesome(s string) int {
	// first[mask] = earliest index where this mask first appeared
	first := make(map[int]int)
	first[0] = -1 // empty prefix has mask 0

	mask := 0
	maxLen := 1

	for i, ch := range s {
		bit := 1 << (ch - '0')
		mask ^= bit

		// Check same mask (all even frequencies)
		if idx, ok := first[mask]; ok {
			maxLen = max(maxLen, i-idx)
		}

		// Check masks differing by one bit (one odd frequency)
		for d := 0; d < 10; d++ {
			neighbor := mask ^ (1 << d)
			if idx, ok := first[neighbor]; ok {
				maxLen = max(maxLen, i-idx)
			}
		}

		// Store first occurrence of this mask
		if _, ok := first[mask]; !ok {
			first[mask] = i
		}
	}

	return maxLen
}
