package main

// LeetCode #856: Score of Parentheses
// https://leetcode.com/problems/score-of-parentheses/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreOfParentheses("()"))
	fmt.Println(ScoreOfParentheses("(())"))
	fmt.Println(ScoreOfParentheses("()()"))
	fmt.Println(ScoreOfParentheses("(()(()))"))
}

// Time: O(n) | Space: O(n)
func ScoreOfParentheses(s string) int {
	stack := []int{0}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, 0)
		} else {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if x != 0 {
				x *= 2
			} else {
				x = 1
			}
			stack[len(stack)-1] += x
		}
	}
	return stack[0]
}
