package main

// LeetCode #3438: Find Valid Pair of Adjacent Digits in String
// https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindValidPairOfAdjacentDigitsInString("2523533"))
	fmt.Println(FindValidPairOfAdjacentDigitsInString("111"))
}

// FindValidPairOfAdjacentDigitsInString finds the first pair of adjacent equal digits where the digit's frequency > the digit.
// Time: O(n). Space: O(n).
func FindValidPairOfAdjacentDigitsInString(s string) string {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt := freq[s[i]]
			if cnt > int(s[i]-'0') {
				return s[i : i+2]
			}
		}
	}
	return ""
}
