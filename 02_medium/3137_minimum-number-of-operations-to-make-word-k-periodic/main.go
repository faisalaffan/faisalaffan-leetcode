package main

// LeetCode #3137: Minimum Number of Operations to Make Word K-Periodic
// https://leetcode.com/problems/minimum-number-of-operations-to-make-word-k-periodic/
// Difficulty: Medium
// Time: O(n) | Space: O(n / k)

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	freq := make(map[string]int)
	maxFreq := 0

	for i := 0; i < n; i += k {
		sub := word[i : i+k]
		freq[sub]++
		if freq[sub] > maxFreq {
			maxFreq = freq[sub]
		}
	}

	return n/k - maxFreq
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4)) // Expected: 1
	fmt.Println(minimumOperationsToMakeKPeriodic("abcabcabc", 3))    // Expected: 0
	fmt.Println(minimumOperationsToMakeKPeriodic("aabbccddee", 5))   // Expected: 1
}
