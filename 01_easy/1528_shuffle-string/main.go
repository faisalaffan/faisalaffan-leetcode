package main

// LeetCode #1528: Shuffle String
// https://leetcode.com/problems/shuffle-string/
// Difficulty: Easy
//
// LeetCode submission: func restoreString(s string, indices []int) string

import "fmt"

func main() {
	fmt.Println(ShuffleString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) // "leetcode"
	fmt.Println(ShuffleString("abc", []int{0, 1, 2}))                     // "abc"
}

// Time: O(n), Space: O(n)
func ShuffleString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, idx := range indices {
		res[idx] = s[i]
	}
	return string(res)
}
