package main

// LeetCode #119: Pascal's Triangle II
// https://leetcode.com/problems/pascals-triangle-ii/
// Difficulty: Easy

import "fmt"

// Time: O(rowIndex^2) | Space: O(rowIndex)
func GetRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func main() {
	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(4))
}
