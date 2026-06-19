package main

// LeetCode #777: Swap Adjacent in LR String
// https://leetcode.com/problems/swap-adjacent-in-lr-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("X", "L"))
}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	n := len(start)
	i, j := 0, 0

	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}

		if i == n && j == n {
			return true
		}
		if i == n || j == n {
			return false
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}

		i++
		j++
	}

	return true
}
