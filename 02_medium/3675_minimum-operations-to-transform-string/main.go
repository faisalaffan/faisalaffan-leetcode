package main

// LeetCode #3675: Minimum Operations to Transform String
// https://leetcode.com/problems/minimum-operations-to-transform-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformString(s string) int {
	minChar := byte('z' + 1)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c < minChar {
			minChar = c
			if minChar == 'b' {
				break
			}
		}
	}
	if minChar > 'z' {
		return 0
	}
	return int('z' + 1 - minChar)
}

func main() {
	fmt.Println(minimumOperationsToTransformString("yz"))
	fmt.Println(minimumOperationsToTransformString("a"))
	fmt.Println(minimumOperationsToTransformString("abc"))
}
