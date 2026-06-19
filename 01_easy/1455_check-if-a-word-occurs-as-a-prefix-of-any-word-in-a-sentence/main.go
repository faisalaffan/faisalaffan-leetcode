package main

// LeetCode #1455: Check If a Word Occurs As a Prefix of Any Word in a Sentence
// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// Difficulty: Easy
//
// LeetCode submission: func isPrefixOfWord(sentence string, searchWord string) int

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("i love eating burger", "burg")) // 4
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("this problem is an easy problem", "pro")) // 2
	fmt.Println(CheckIfAWordOccursAsAPrefixOfAnyWordInASentence("i am tired", "you")) // -1
}

// Time: O(n), Space: O(n)
func CheckIfAWordOccursAsAPrefixOfAnyWordInASentence(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		if strings.HasPrefix(w, searchWord) {
			return i + 1
		}
	}
	return -1
}
