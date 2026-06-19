package main

// LeetCode #2785: Sort Vowels in a String
// https://leetcode.com/problems/sort-vowels-in-a-string/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func SortVowelsInAString(s string) string {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	vowels := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	result := []byte(s)
	vi := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = vowels[vi]
			vi++
		}
	}

	return string(result)
}

func main() {
	fmt.Println(SortVowelsInAString("lEetcOde"))
	fmt.Println(SortVowelsInAString("lYmpH"))
}
