package main

// LeetCode #1832: Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckIfPangram(sentence string) bool {
	seen := 0
	for i := 0; i < len(sentence); i++ {
		seen |= 1 << (sentence[i] - 'a')
	}
	return seen == (1<<26)-1
}

func main() {
	fmt.Println(CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(CheckIfPangram("leetcode"))
}
