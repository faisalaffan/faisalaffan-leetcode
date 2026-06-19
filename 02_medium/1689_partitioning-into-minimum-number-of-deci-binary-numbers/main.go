package main

// LeetCode #1689: Partitioning Into Minimum Number Of Deci-Binary Numbers
// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minPartitions(n string) int {
	maxDigit := 0
	for _, ch := range n {
		digit := int(ch - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
		if maxDigit == 9 {
			break // can't get higher than 9
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(minPartitions("32"))     // Expected: 3
	fmt.Println(minPartitions("82734"))  // Expected: 8
	fmt.Println(minPartitions("27346209830709182346")) // Expected: 9
}
