package main

// LeetCode #1668: Maximum Repeating Substring
// https://leetcode.com/problems/maximum-repeating-substring/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n * m), Space: O(n) where n = len(sequence), m = len(word)
func MaxRepeating(sequence string, word string) int {
	k := 0
	repeated := word
	for strings.Contains(sequence, repeated) {
		k++
		repeated += word
	}
	return k
}

func main() {
	fmt.Println(MaxRepeating("ababc", "ab"))
	fmt.Println(MaxRepeating("ababc", "ba"))
	fmt.Println(MaxRepeating("ababc", "ac"))
}
