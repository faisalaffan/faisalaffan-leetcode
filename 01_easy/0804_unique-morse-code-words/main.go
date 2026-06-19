package main

// LeetCode #804: Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/
// Difficulty: Easy

import "fmt"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) // 2
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))                        // 1
}

// uniqueMorseRepresentations counts unique Morse code transformations of words.
// Time: O(n * m) where n = len(words), m = avg len. Space: O(n).
func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool)
	for _, word := range words {
		var code string
		for _, c := range word {
			code += morse[c-'a']
		}
		set[code] = true
	}
	return len(set)
}
