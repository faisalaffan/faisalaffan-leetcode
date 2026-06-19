package main

// LeetCode #2703: Return Length of Arguments Passed
// https://leetcode.com/problems/return-length-of-arguments-passed/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns number of arguments.

import "fmt"

func main() {
	fmt.Println(ReturnLengthOfArgumentsPassed(1, 2, 3))
	fmt.Println(ReturnLengthOfArgumentsPassed("a", "b"))
}

func ReturnLengthOfArgumentsPassed(args ...interface{}) int {
	return len(args)
}
