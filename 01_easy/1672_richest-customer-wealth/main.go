package main

// LeetCode #1672: Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
// Difficulty: Easy

import "fmt"

// Time: O(m*n), Space: O(1)
func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		sum := 0
		for _, amount := range customer {
			sum += amount
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}

func main() {
	fmt.Println(MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(MaximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(MaximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
