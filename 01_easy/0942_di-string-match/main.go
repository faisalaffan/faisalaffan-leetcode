package main

// LeetCode #942: DI String Match
// https://leetcode.com/problems/di-string-match/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID")) // [0,4,1,3,2]
	fmt.Println(diStringMatch("III"))  // [0,1,2,3]
	fmt.Println(diStringMatch("DDI"))  // [3,2,0,1]
}

// diStringMatch returns a permutation that matches the DI pattern.
// Time: O(n). Space: O(n).
func diStringMatch(s string) []int {
	n := len(s)
	result := make([]int, n+1)
	low, high := 0, n
	for i, c := range s {
		if c == 'I' {
			result[i] = low
			low++
		} else {
			result[i] = high
			high--
		}
	}
	result[n] = low // or high (they are equal)
	return result
}
