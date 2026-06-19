package main

// LeetCode #159: Longest Substring with At Most Two Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	charCount := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		for len(charCount) > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abc"))
}
