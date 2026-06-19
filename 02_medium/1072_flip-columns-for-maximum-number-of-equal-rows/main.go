package main

// LeetCode #1072: Flip Columns For Maximum Number of Equal Rows
// https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/
// Difficulty: Medium
//
// Approach: Normalize each row to a pattern (starting with 0).
//           Rows with same or complementary pattern can be made equal.
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 1}, {1, 1}})) // 1
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}})) // 2
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patternCount := make(map[string]int)

	for _, row := range matrix {
		pattern := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			if row[0] == 0 {
				pattern[j] = byte('0' + row[j])
			} else {
				pattern[j] = byte('0' + 1 - row[j])
			}
		}
		patternCount[string(pattern)]++
	}

	result := 0
	for _, count := range patternCount {
		if count > result {
			result = count
		}
	}
	return result
}
