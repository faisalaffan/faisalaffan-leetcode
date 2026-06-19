package main

// LeetCode #3368: First Letter Capitalization
// https://leetcode.com/problems/first-letter-capitalization/
// Difficulty: Hard [Paid]
//
// Capitalize first letter of each word, lowercase rest.

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(FirstLetterCapitalization("hello world"))
	fmt.Println(FirstLetterCapitalization("Leetcode is fun"))
}

func FirstLetterCapitalization(title string) string {
	words := strings.Fields(title)
	for i, w := range words {
		r := []rune(w)
		for j := range r {
			if j == 0 {
				r[j] = unicode.ToUpper(r[j])
			} else {
				r[j] = unicode.ToLower(r[j])
			}
		}
		words[i] = string(r)
	}
	return strings.Join(words, " ")
}
