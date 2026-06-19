package main

// LeetCode #3070: Count Submatrices with Top-Left Element and Sum Less Than k
// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20))
	fmt.Println(countSubmatrices([][]int{{1, 2}, {3, 4}}, 5))
}

func countSubmatrices(grid [][]int, k int) (ans int) {
	m, n := len(grid), len(grid[0])
	pref := make([][]int, m+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pref[i+1][j+1] = grid[i][j] + pref[i][j+1] + pref[i+1][j] - pref[i][j]
			if pref[i+1][j+1] <= k {
				ans++
			}
		}
	}
	return
}
