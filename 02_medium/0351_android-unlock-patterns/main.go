package main

// LeetCode #351: Android Unlock Patterns
// https://leetcode.com/problems/android-unlock-patterns/
// Difficulty: Medium [Paid]
// Time: O(n!) | Space: O(n)

import "fmt"

func numberOfPatterns(m int, n int) int {
	// skip[i][j] = key that must be visited between i and j (0 if none)
	skip := [10][10]int{}
	skip[1][3] = 2
	skip[3][1] = 2
	skip[1][7] = 4
	skip[7][1] = 4
	skip[3][9] = 6
	skip[9][3] = 6
	skip[7][9] = 8
	skip[9][7] = 8
	skip[1][9] = 5
	skip[9][1] = 5
	skip[2][8] = 5
	skip[8][2] = 5
	skip[3][7] = 5
	skip[7][3] = 5
	skip[4][6] = 5
	skip[6][4] = 5

	visited := [10]bool{}
	var dfs func(cur int, remaining int) int
	dfs = func(cur int, remaining int) int {
		if remaining == 0 {
			return 1
		}
		visited[cur] = true
		count := 0
		for next := 1; next <= 9; next++ {
			if !visited[next] && (skip[cur][next] == 0 || visited[skip[cur][next]]) {
				count += dfs(next, remaining-1)
			}
		}
		visited[cur] = false
		return count
	}

	total := 0
	for length := m; length <= n; length++ {
		// Start from 1, 2, 5 (symmetry: 1,3,7,9 are same; 2,4,6,8 are same)
		total += dfs(1, length-1) * 4
		total += dfs(2, length-1) * 4
		total += dfs(5, length-1)
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfPatterns(1, 1))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", numberOfPatterns(1, 2))
	// Expected: 65

	// Test case 3
	fmt.Println("Test 3:", numberOfPatterns(3, 3))
	// Expected: 320 (length exactly 3)
}
