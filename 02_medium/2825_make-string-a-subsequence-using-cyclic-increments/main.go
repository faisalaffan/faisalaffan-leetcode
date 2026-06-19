package main

// LeetCode #2825: Make String a Subsequence Using Cyclic Increments
// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool {
	j := 0
	for i := 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || (str1[i]-'a'+1)%26 == str2[j]-'a' {
			j++
		}
	}
	return j == len(str2)
}

func main() {
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "bcd"))
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "ad"))
}
