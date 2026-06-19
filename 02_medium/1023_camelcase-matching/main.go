package main

// LeetCode #1023: Camelcase Matching
// https://leetcode.com/problems/camelcase-matching/
// Difficulty: Medium
//
// Approach: For each query, use two pointers to match pattern characters
// Time: O(n * (len(query) + len(pattern)))
// Space: O(n) for output

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	// [true,false,true,true,false]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	// [true,false,true,false,false]
}

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))

	for i, q := range queries {
		result[i] = matches(q, pattern)
	}

	return result
}

func matches(query, pattern string) bool {
	j := 0
	for i := 0; i < len(query); i++ {
		if j < len(pattern) && query[i] == pattern[j] {
			j++
		} else if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		}
	}
	return j == len(pattern)
}
