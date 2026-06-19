package main

// LeetCode #2973: Find Number of Coins to Place in Tree Nodes
// https://leetcode.com/problems/find-number-of-coins-to-place-in-tree-nodes/
// Difficulty: Hard
//
// For each subtree, find the maximum product of 3 values (or 0 if
// fewer than 3 nodes in subtree). The product can be:
//   - 3 largest positive values
//   - 2 smallest negative values * largest positive value
//
// DFS returns the up to 5 most relevant values (2 smallest, 3 largest)
// for computing the product.

import (
	"fmt"
	"sort"
)

func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	ans := make([]int64, n)
	for i := range ans {
		ans[i] = 1
	}

	var dfs func(a, fa int) []int
	dfs = func(a, fa int) []int {
		res := []int{cost[a]}
		for _, b := range g[a] {
			if b != fa {
				res = append(res, dfs(b, a)...)
			}
		}
		sort.Ints(res)
		m := len(res)

		if m >= 3 {
			// Option 1: three largest
			x := res[m-1] * res[m-2] * res[m-3]
			// Option 2: two smallest (most negative) * largest
			y := res[0] * res[1] * res[m-1]
			if x > y {
				y = x
			}
			if y > 0 {
				ans[a] = int64(y)
			} else {
				ans[a] = 0
			}
		}

		// Keep at most 5 values: 2 smallest + 3 largest
		if m >= 5 {
			res = append(res[:2], res[m-3:]...)
		}
		return res
	}

	dfs(0, -1)
	return ans
}

func main() {
	// Example: tree with 5 nodes
	fmt.Println(placedCoins([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, []int{1, 10, 1, 1, 1}))
	fmt.Println(placedCoins([][]int{{0, 1}, {1, 2}}, []int{1, 2, 3}))
	fmt.Println(placedCoins([][]int{{0, 1}}, []int{5, 5}))
}
