package main

// LeetCode #3258: Count Substrings That Satisfy K-Constraint I
// https://leetcode.com/problems/count-substrings-that-satisfy-k-constraint-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("10101", 1))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("1010101", 2))
	fmt.Println(CountSubstringsThatSatisfyKConstraintI("11111", 1))
}

// CountSubstringsThatSatisfyKConstraintI counts substrings where both number of 0s and 1s <= k.
// Time: O(n^2). Space: O(1).
func CountSubstringsThatSatisfyKConstraintI(s string, k int) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		zeros, ones := 0, 0
		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeros++
			} else {
				ones++
			}
			if zeros <= k || ones <= k {
				count++
			} else {
				break
			}
		}
	}
	return count
}
