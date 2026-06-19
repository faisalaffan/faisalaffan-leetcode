package main

// LeetCode #823: Binary Trees With Factors
// https://leetcode.com/problems/binary-trees-with-factors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BinaryTreesWithFactors([]int{2, 4}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 4, 5, 10}))
	fmt.Println(BinaryTreesWithFactors([]int{2, 3, 4, 6, 8, 12, 24}))
}

// Time: O(n^2) | Space: O(n)
func BinaryTreesWithFactors(arr []int) int {
	const mod = 1_000_000_007
	sort.Ints(arr)

	dp := make(map[int]int)
	for _, x := range arr {
		dp[x] = 1
	}

	for i, x := range arr {
		for j := 0; j < i; j++ {
			if x%arr[j] == 0 {
				if val, ok := dp[x/arr[j]]; ok {
					dp[x] = (dp[x] + dp[arr[j]]*val) % mod
				}
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return ans
}
