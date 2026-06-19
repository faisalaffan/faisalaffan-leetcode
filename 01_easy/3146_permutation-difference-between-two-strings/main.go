package main

// LeetCode #3146: Permutation Difference between Two Strings
// https://leetcode.com/problems/permutation-difference-between-two-strings/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPermutationDifference
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abc", "bac")) // 2
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abcde", "edcba")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: findPermutationDifference
func PermutationDifferenceBetweenTwoStrings(s string, t string) int {
	pos := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		pos[t[i]] = i
	}
	diff := 0
	for i := 0; i < len(s); i++ {
		d := i - pos[s[i]]
		if d < 0 {
			d = -d
		}
		diff += d
	}
	return diff
}
