package main

// LeetCode #1544: Make The String Great
// https://leetcode.com/problems/make-the-string-great/
// Difficulty: Easy
//
// LeetCode submission: func makeGood(s string) string

import "fmt"

func main() {
	fmt.Println(MakeTheStringGreat("leEeetcode")) // "leetcode"
	fmt.Println(MakeTheStringGreat("abBAcC"))     // ""
	fmt.Println(MakeTheStringGreat("s"))          // "s"
}

// Time: O(n), Space: O(n)
func MakeTheStringGreat(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		stack = append(stack, s[i])
		n := len(stack)
		if n >= 2 {
			diff := int(stack[n-1]) - int(stack[n-2])
			if diff == 32 || diff == -32 {
				stack = stack[:n-2]
			}
		}
	}
	return string(stack)
}
