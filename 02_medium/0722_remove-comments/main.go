package main

// LeetCode #722: Remove Comments
// https://leetcode.com/problems/remove-comments/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	source := []string{
		"/*Test program */",
		"int main()",
		"{ ",
		"  // variable declaration ",
		"int a, b, c;",
		"/* This is a test",
		"   multiline  ",
		"   comment for ",
		"   testing */",
		"a = b + c;",
		"}",
	}
	result := removeComments(source)
	for _, line := range result {
		fmt.Println(line)
	}
}

func removeComments(source []string) []string {
	result := make([]string, 0)
	inBlock := false
	var current strings.Builder

	for _, line := range source {
		i := 0
		n := len(line)

		if !inBlock {
			current.Reset()
		}

		for i < n {
			if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '*' {
				inBlock = true
				i += 2
			} else if inBlock && i+1 < n && line[i] == '*' && line[i+1] == '/' {
				inBlock = false
				i += 2
			} else if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '/' {
				break
			} else if !inBlock {
				current.WriteByte(line[i])
				i++
			} else {
				i++
			}
		}

		if !inBlock && current.Len() > 0 {
			result = append(result, current.String())
		}
	}

	return result
}
