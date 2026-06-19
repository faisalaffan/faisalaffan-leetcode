package main

// LeetCode #828: Count Unique Characters of All Substrings of a Given String
// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/
// Difficulty: Hard
// Approach: Contribution per character. For each s[i], count substrings where s[i] is the
// first occurrence of that character within the substring. Use prev/next occurrence arrays.

import "fmt"

func uniqueLetterString(s string) int {
	n := len(s)
	prev := make([]int, n)
	next := make([]int, n)
	last := make([]int, 26)

	for i := range last {
		last[i] = -1
	}
	for i := 0; i < n; i++ {
		c := int(s[i] - 'A')
		prev[i] = last[c]
		last[c] = i
	}
	for i := range last {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		c := int(s[i] - 'A')
		next[i] = last[c]
		last[c] = i
	}

	ans := 0
	for i := 0; i < n; i++ {
		left := i - prev[i]
		right := next[i] - i
		ans += left * right
	}
	return ans
}

func main() {
	fmt.Println(uniqueLetterString("ABC")) // Expected: 10
	fmt.Println(uniqueLetterString("ABA")) // Expected: 8
	fmt.Println(uniqueLetterString("LEETCODE")) // Additional test
}
