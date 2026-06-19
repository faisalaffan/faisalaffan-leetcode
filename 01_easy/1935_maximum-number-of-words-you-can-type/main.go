package main

// LeetCode #1935: Maximum Number of Words You Can Type
// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsYouCanType("hello world", "ad"))                   // 1
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "e"))                      // 0
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "lt"))                     // 1
}

// Time: O(n + m), Space: O(k) where k = len(brokenLetters)
func MaximumNumberOfWordsYouCanType(text string, brokenLetters string) int {
	broken := make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		broken[brokenLetters[i]] = true
	}

	words := strings.Fields(text)
	count := 0
	for _, word := range words {
		canType := true
		for i := 0; i < len(word); i++ {
			if broken[word[i]] {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}
