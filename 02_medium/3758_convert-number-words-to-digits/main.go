package main

// LeetCode #3758: Convert Number Words to Digits
// https://leetcode.com/problems/convert-number-words-to-digits/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strings"
)

func convertNumberWordsToDigits(s string) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var res strings.Builder
	i := 0
	for i < len(s) {
		found := false
		for d, word := range words {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				res.WriteByte(byte(d) + '0')
				i += len(word)
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return res.String()
}

func main() {
	fmt.Println(convertNumberWordsToDigits("onefourthree"))
	fmt.Println(convertNumberWordsToDigits("ninexsix"))
	fmt.Println(convertNumberWordsToDigits("zeero"))
}
