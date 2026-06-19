package main

// LeetCode #3561: Resulting String After Adjacent Removals
// https://leetcode.com/problems/resulting-string-after-adjacent-removals/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ResultingStringAfterAdjacentRemovals("abbaca"))
	// Test case 2
	fmt.Println("Test 2:", ResultingStringAfterAdjacentRemovals("azxxzy"))
	// Test case 3
	fmt.Println("Test 3:", ResultingStringAfterAdjacentRemovals("a"))
}

func ResultingStringAfterAdjacentRemovals(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
