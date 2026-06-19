package main

// LeetCode #2168: Unique Substrings With Equal Digit Frequency
// https://leetcode.com/problems/unique-substrings-with-equal-digit-frequency/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func equalDigitFrequency(s string) int {
	n := len(s)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
		freq := make([]int, 10)
		distinct := 0
		maxFreq := 0
		for j := i; j < n; j++ {
			d := int(s[j] - '0')
			if freq[d] == 0 {
				distinct++
			}
			freq[d]++
			if freq[d] > maxFreq {
				maxFreq = freq[d]
			}
			// All digits appear same frequency iff distinct * maxFreq == totalLen
			if distinct*maxFreq == j-i+1 {
				seen[s[i:j+1]] = true
			}
		}
	}

	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalDigitFrequency("1212"))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", equalDigitFrequency("12321"))
	// Expected: 9
}
