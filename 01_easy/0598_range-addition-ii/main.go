package main

// LeetCode #598: Range Addition II
// https://leetcode.com/problems/range-addition-ii/
// Difficulty: Easy

import "fmt"

// Time: O(k), Space: O(1)
func RangeAdditionIi(m, n int, ops [][]int) int {
	minA, minB := m, n
	for _, op := range ops {
		if op[0] < minA {
			minA = op[0]
		}
		if op[1] < minB {
			minB = op[1]
		}
	}
	return minA * minB
}

func main() {
	fmt.Println(RangeAdditionIi(3, 3, [][]int{{2, 2}, {3, 3}}))
	fmt.Println(RangeAdditionIi(3, 3, [][]int{}))
}
