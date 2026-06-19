package main

// LeetCode #1678: Goal Parser Interpretation
// https://leetcode.com/problems/goal-parser-interpretation/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func Interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

func main() {
	fmt.Println(Interpret("G()(al)"))
	fmt.Println(Interpret("G()()()()(al)"))
	fmt.Println(Interpret("(al)G(al)()()G"))
}
