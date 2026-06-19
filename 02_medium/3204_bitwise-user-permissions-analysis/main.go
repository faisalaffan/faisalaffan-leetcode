package main

// LeetCode #3204: Bitwise User Permissions Analysis
// https://leetcode.com/problems/bitwise-user-permissions-analysis/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func bitwiseUserPermissions(permissions [][]int) int {
	if len(permissions) == 0 {
		return 0
	}

	combined := permissions[0][1]
	for i := 1; i < len(permissions); i++ {
		combined |= permissions[i][1]
	}
	return combined
}

func main() {
	fmt.Println(bitwiseUserPermissions([][]int{{1, 1}, {2, 2}, {3, 4}})) // Expected: 7 (1|2|4)
	fmt.Println(bitwiseUserPermissions([][]int{{1, 8}, {2, 3}}))         // Expected: 11 (8|3)
}
