package main

// LeetCode #2944: Minimum Number of Coins for Fruits
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))
	fmt.Println(minimumCoins([]int{26, 53, 10, 24, 25, 20, 63, 51}))
}

func minimumCoins(prices []int) int {
	n := len(prices)
	f := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}
		if f[i] == 0 {
			f[i] = 1 << 30
			for j := i + 1; j <= i*2+1; j++ {
				cost := dfs(j) + prices[i-1]
				if cost < f[i] {
					f[i] = cost
				}
			}
		}
		return f[i]
	}
	return dfs(1)
}
