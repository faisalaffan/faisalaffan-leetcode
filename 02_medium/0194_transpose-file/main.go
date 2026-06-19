package main

// LeetCode #194: Transpose File
// https://leetcode.com/problems/transpose-file/
// Difficulty: Medium
// Time: O(m*n), Space: O(m*n)

import (
	"fmt"
	"strings"
)

func transposeFile(content string) []string {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) == 0 || lines[0] == "" {
		return nil
	}

	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Fields(line)
	}

	if len(matrix) == 0 {
		return nil
	}

	cols := 0
	for _, row := range matrix {
		if len(row) > cols {
			cols = len(row)
		}
	}

	result := make([]string, cols)
	for c := 0; c < cols; c++ {
		var row []string
		for r := 0; r < len(matrix); r++ {
			if c < len(matrix[r]) {
				row = append(row, matrix[r][c])
			}
		}
		result[c] = strings.Join(row, " ")
	}

	return result
}

func main() {
	content := "name age\nalice 21\nryan 30"
	for _, line := range transposeFile(content) {
		fmt.Println(line)
	}
	fmt.Println("---")
	content2 := "a b c\nd e f"
	for _, line := range transposeFile(content2) {
		fmt.Println(line)
	}
}
