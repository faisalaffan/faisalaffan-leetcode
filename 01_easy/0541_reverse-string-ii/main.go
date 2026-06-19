package main

// LeetCode #541: Reverse String II
// https://leetcode.com/problems/reverse-string-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseStringIi(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		lo, hi := i, i+k-1
		if hi >= len(b) {
			hi = len(b) - 1
		}
		for lo < hi {
			b[lo], b[hi] = b[hi], b[lo]
			lo++
			hi--
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseStringIi("abcdefg", 2))
	fmt.Println(ReverseStringIi("abcd", 2))
}
