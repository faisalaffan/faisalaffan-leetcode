package main

// LeetCode #1614: Maximum Nesting Depth of the Parentheses
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaximumNestingDepthOfTheParentheses(s string) int {
	maxDepth, currentDepth := 0, 0
	for _, ch := range s {
		if ch == '(' {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		} else if ch == ')' {
			currentDepth--
		}
	}
	return maxDepth
}

func main() {
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1+(2*3)+((8)/4))+1"))
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1)+((2))+(((3)))"))
	fmt.Println(MaximumNestingDepthOfTheParentheses(""))
}
