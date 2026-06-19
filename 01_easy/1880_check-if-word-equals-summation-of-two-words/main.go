package main

// LeetCode #1880: Check if Word Equals Summation of Two Words
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordValue(firstWord)+wordValue(secondWord) == wordValue(targetWord)
}

func wordValue(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'a')
	}
	return val
}

func main() {
	fmt.Println(IsSumEqual("acb", "cba", "cdb"))
	fmt.Println(IsSumEqual("aaa", "a", "aab"))
	fmt.Println(IsSumEqual("aaa", "a", "aaaa"))
}
