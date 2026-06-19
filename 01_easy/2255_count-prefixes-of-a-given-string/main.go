package main

// LeetCode #2255: Count Prefixes of a Given String
// https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))           // 3
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "a"}, "aa"))                                      // 2
	fmt.Println(CountPrefixesOfAGivenString([]string{"feh", "w", "w", "l", "w", "o", "w", "o", "w"}, "w")) // 0
}

func CountPrefixesOfAGivenString(words []string, s string) int {
	count := 0
	for _, w := range words {
		if len(w) <= len(s) && s[:len(w)] == w {
			count++
		}
	}
	return count
}
