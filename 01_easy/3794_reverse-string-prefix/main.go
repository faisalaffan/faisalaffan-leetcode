package main

// LeetCode #3794: Reverse String Prefix
// https://leetcode.com/problems/reverse-string-prefix/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseStringPrefix("abcd", 2))
	fmt.Println(ReverseStringPrefix("xyz", 3))
	fmt.Println(ReverseStringPrefix("hey", 1))
}

// Time: O(n)
// Space: O(n)
func ReverseStringPrefix(s string, k int) string {
	b := []byte(s)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
