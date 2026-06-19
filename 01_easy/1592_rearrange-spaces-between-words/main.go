package main

// LeetCode #1592: Rearrange Spaces Between Words
// https://leetcode.com/problems/rearrange-spaces-between-words/
// Difficulty: Easy
//
// LeetCode submission: func reorderSpaces(text string) string

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RearrangeSpacesBetweenWords("  this   is  a sentence ")) // "this   is   a   sentence"
	fmt.Println(RearrangeSpacesBetweenWords(" practice   makes   perfect")) // "practice   makes   perfect "
}

// Time: O(n), Space: O(n)
func RearrangeSpacesBetweenWords(text string) string {
	words := strings.Fields(text)
	spaces := strings.Count(text, " ")
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaces)
	}
	between := spaces / (len(words) - 1)
	extra := spaces % (len(words) - 1)
	res := strings.Join(words, strings.Repeat(" ", between))
	res += strings.Repeat(" ", extra)
	return res
}
