package main

// LeetCode #3619: Count Islands With Total Value Divisible by K
// https://leetcode.com/problems/count-islands-with-total-value-divisible-by-k/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Println("Test 1:", CountIslandsWithTotalValueDivisibleByK(grid, 2))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 2:", CountIslandsWithTotalValueDivisibleByK(grid2, 4))
	// Test case 3
	grid3 := [][]int{{0, 0}, {0, 0}}
	fmt.Println("Test 3:", CountIslandsWithTotalValueDivisibleByK(grid3, 1))
}

func CountIslandsWithTotalValueDivisibleByK(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || grid[r][c] == 0 {
			return 0
		}
		visited[r][c] = true
		sum := grid[r][c]
		sum += dfs(r-1, c)
		sum += dfs(r+1, c)
		sum += dfs(r, c-1)
		sum += dfs(r, c+1)
		return sum
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] != 0 {
				sum := dfs(i, j)
				if sum%k == 0 {
					count++
				}
			}
		}
	}
	return count
}
