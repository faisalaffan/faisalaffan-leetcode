package main

// LeetCode #1547: Minimum Cost to Cut a Stick
// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
// Difficulty: Hard
//
// DP interval approach:
// 1. Add 0 and n to the cuts array, sort.
// 2. DP[i][j] = minimum cost to cut stick from cuts[i] to cuts[j].
// 3. For each interval [i,j], try every cut point k in (i,j).
// 4. DP[i][j] = min(DP[i][k] + DP[k][j] + (cuts[j]-cuts[i]))
// 5. Return DP[0][len(cuts)-1].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example: n=7, cuts=[1,3,4,5] -> 16
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))

	// Additional tests
	fmt.Println(minCost(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(minCost(3, []int{1}))
	fmt.Println(minCost(4, []int{2}))
}

func minCost(n int, cuts []int) int {
	// Add boundaries and sort
	extended := make([]int, 0, len(cuts)+2)
	extended = append(extended, 0)
	extended = append(extended, cuts...)
	extended = append(extended, n)
	sort.Ints(extended)

	m := len(extended)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// Interval DP: process by increasing length
	for length := 2; length < m; length++ {
		for i := 0; i+length < m; i++ {
			j := i + length
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				cost := dp[i][k] + dp[k][j] + (extended[j] - extended[i])
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	return dp[0][m-1]
}
