package main

// LeetCode #1461: Check If a String Contains All Binary Codes of Size K
// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(hasAllCodes("00110110", 2)) // true

	// Test case 2
	fmt.Println(hasAllCodes("0110", 1)) // true

	// Test case 3
	fmt.Println(hasAllCodes("0110", 2)) // false

	// Test case 4
	fmt.Println(hasAllCodes("000000000010101010111010101001110100110100110010100101001000101010101011", 4)) // true
}

// Time: O(n*k) where n = len(s), or O(n) with bit manipulation
// Space: O(2^k) for storing seen codes
func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	needed := 1 << k
	seen := make([]bool, needed)
	mask := needed - 1
	hash := 0

	// Compute hash for first k bits
	for i := 0; i < k; i++ {
		hash = (hash << 1) | int(s[i]-'0')
	}
	seen[hash] = true
	count := 1

	// Sliding window with rolling hash
	for i := k; i < len(s); i++ {
		hash = ((hash << 1) & mask) | int(s[i]-'0')
		if !seen[hash] {
			seen[hash] = true
			count++
			if count == needed {
				return true
			}
		}
	}

	return count == needed
}
