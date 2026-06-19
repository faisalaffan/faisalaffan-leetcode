package main

// LeetCode #2788: Split Strings by Separator
// https://leetcode.com/problems/split-strings-by-separator/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SplitStringsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
	fmt.Println(SplitStringsBySeparator([]string{"$easy$", "$problem$"}, '$'))
}

func SplitStringsBySeparator(words []string, separator byte) []string {
	result := []string{}
	for _, w := range words {
		parts := strings.Split(w, string(separator))
		for _, p := range parts {
			if p != "" {
				result = append(result, p)
			}
		}
	}
	return result
}
