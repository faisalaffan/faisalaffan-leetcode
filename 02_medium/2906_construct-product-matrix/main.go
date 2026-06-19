package main

// LeetCode #2906: Construct Product Matrix
// https://leetcode.com/problems/construct-product-matrix/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructProductMatrix([][]int{{1, 2}, {3, 4}}))
	fmt.Println(constructProductMatrix([][]int{{2, 3, 4}, {5, 6, 7}}))
}

func constructProductMatrix(grid [][]int) [][]int {
	const mod int = 12345
	n, m := len(grid), len(grid[0])
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, m)
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = p[i][j] * pre % mod
			pre = pre * grid[i][j] % mod
		}
	}
	return p
}
