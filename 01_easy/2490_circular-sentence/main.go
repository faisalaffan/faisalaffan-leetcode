package main

// LeetCode #2490: Circular Sentence
// https://leetcode.com/problems/circular-sentence/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CircularSentence("leetcode exercises sound delightful")) // true
	fmt.Println(CircularSentence("eetcode"))                            // true
	fmt.Println(CircularSentence("Leetcode is cool"))                   // false
}

func CircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
