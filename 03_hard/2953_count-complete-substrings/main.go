package main

// LeetCode #2953: Count Complete Substrings
// https://leetcode.com/problems/count-complete-substrings/
// Difficulty: Hard
//
// Approach: Sliding window with frequency-of-frequencies tracking.
// A substring is complete if:
//   1. Every character appears exactly k times.
//   2. For any two adjacent chars, |ord(c1)-ord(c2)| <= 2.
//
// Step 1: Split string at positions where adjacent diff > 2.
// Step 2: For each valid segment, try window sizes = i*k for i = 1..26.
//   Maintain a frequency array cnt[26] and a freq-of-freq array.
//   Window is complete if freq[k] == i (exactly i chars appear k times).

import (
	"fmt"
)

func countCompleteSubstrings(word string, k int) int {
	n := len(word)

	// Count complete substrings in a segment (where all adjacent diffs <= 2)
	countInSegment := func(s string) int {
		m := len(s)
		res := 0
		// Try each possible distinct character count (1 to 26)
		for distinct := 1; distinct <= 26; distinct++ {
			winLen := distinct * k
			if winLen > m {
				break
			}

			// Frequency of each character in the current window
			cnt := make([]int, 26)
			// freq[count] = how many characters appear exactly `count` times
			freq := make([]int, m+1)
			freq[0] = 26

			// Initialize first window
			for i := 0; i < winLen; i++ {
				idx := s[i] - 'a'
				freq[cnt[idx]]--
				cnt[idx]++
				freq[cnt[idx]]++
			}
			if freq[k] == distinct {
				res++
			}

			// Slide the window
			for i := winLen; i < m; i++ {
				// Add right character
				right := s[i] - 'a'
				freq[cnt[right]]--
				cnt[right]++
				freq[cnt[right]]++

				// Remove left character
				left := s[i-winLen] - 'a'
				freq[cnt[left]]--
				cnt[left]--
				freq[cnt[left]]++

				if freq[k] == distinct {
					res++
				}
			}
		}
		return res
	}

	ans := 0
	start := 0
	for i := 1; i < n; i++ {
		diff := int(word[i]) - int(word[i-1])
		if diff < 0 {
			diff = -diff
		}
		if diff > 2 {
			ans += countInSegment(word[start:i])
			start = i
		}
	}
	ans += countInSegment(word[start:])
	return ans
}

func main() {
	// Example: "igigee", k=2 -> 3
	// Complete substrings: "ig", "gi", "ee" (each with 2 occurrences of 1 char)
	fmt.Println(countCompleteSubstrings("igigee", 2))

	// Simple cases
	fmt.Println(countCompleteSubstrings("aa", 2))
	fmt.Println(countCompleteSubstrings("abc", 1))
}
