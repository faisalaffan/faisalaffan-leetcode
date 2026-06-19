package main

// LeetCode #1119: Remove Vowels from a String
// https://leetcode.com/problems/remove-vowels-from-a-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
	fmt.Println(removeVowels("aeiou"))                         // ""
}

// LeetCode submission: removeVowels
func removeVowels(s string) string {
	ans := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
