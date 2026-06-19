package main

// LeetCode #1957: Delete Characters to Make Fancy String
// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DeleteCharactersToMakeFancyString("leeetcode"))     // "leetcode"
	fmt.Println(DeleteCharactersToMakeFancyString("aaabaaaa"))      // "aabaa"
	fmt.Println(DeleteCharactersToMakeFancyString("aab"))           // "aab"
}

// Time: O(n), Space: O(n)
func DeleteCharactersToMakeFancyString(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n >= 2 && result[n-1] == s[i] && result[n-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
