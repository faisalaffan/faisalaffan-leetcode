package main

// LeetCode #2000: Reverse Prefix of Word
// https://leetcode.com/problems/reverse-prefix-of-word/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReversePrefixOfWord("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(ReversePrefixOfWord("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(ReversePrefixOfWord("abcd", 'z'))    // "abcd"
}

// Time: O(n), Space: O(n)
func ReversePrefixOfWord(word string, ch byte) string {
	idx := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}

	result := make([]byte, len(word))
	for i := 0; i <= idx; i++ {
		result[i] = word[idx-i]
	}
	for i := idx + 1; i < len(word); i++ {
		result[i] = word[i]
	}
	return string(result)
}
