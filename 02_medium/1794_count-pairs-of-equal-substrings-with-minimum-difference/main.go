package main

// LeetCode #1794: Count Pairs of Equal Substrings With Minimum Difference
// https://leetcode.com/problems/count-pairs-of-equal-substrings-with-minimum-difference/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func countQuadruples(firstString string, secondString string) int {
	firstPos := make([]int, 26)
	lastPos := make([]int, 26)
	for i := range firstPos {
		firstPos[i] = -1
		lastPos[i] = -1
	}

	for i, ch := range firstString {
		idx := ch - 'a'
		if firstPos[idx] == -1 {
			firstPos[idx] = i
		}
	}
	for i, ch := range secondString {
		idx := ch - 'a'
		lastPos[idx] = i
	}

	minDiff := 1 << 30
	count := 0

	for i := 0; i < 26; i++ {
		if firstPos[i] != -1 && lastPos[i] != -1 {
			diff := firstPos[i] - lastPos[i]
			if diff < minDiff {
				minDiff = diff
				count = 1
			} else if diff == minDiff {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countQuadruples("abcd", "bcd")) // test 1
	fmt.Println(countQuadruples("abc", "abc")) // test 2
	fmt.Println(countQuadruples("abb", "b")) // test 3
}
