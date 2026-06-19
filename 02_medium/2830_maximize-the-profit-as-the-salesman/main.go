package main

// LeetCode #2830: Maximize the Profit as the Salesman
// https://leetcode.com/problems/maximize-the-profit-as-the-salesman/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func MaximizeTheProfitAsTheSalesman(n int, offers [][]int) int {
	// Group offers by end position
	byEnd := make([][][]int, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		byEnd[end] = append(byEnd[end], []int{start, gold})
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		for _, offer := range byEnd[i] {
			start, gold := offer[0], offer[1]
			val := gold
			if start > 0 {
				val += dp[start-1]
			}
			if val > dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(MaximizeTheProfitAsTheSalesman(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
	fmt.Println(MaximizeTheProfitAsTheSalesman(3, [][]int{{0, 0, 5}, {1, 1, 3}, {2, 2, 4}}))
}
