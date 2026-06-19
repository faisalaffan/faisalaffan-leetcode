package main

// LeetCode #3004: Maximum Subtree of the Same Color (PAID)
// https://leetcode.com/problems/maximum-subtree-of-the-same-color/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

// Given an undirected tree rooted at 0, find the size of the largest subtree
// where all nodes have the same color. A subtree is rooted at some node v
// and includes all of v's descendants.

import "fmt"

func main() {
	// Test 1: Entire tree is same color
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}
	colors := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors)) // 7

	// Test 2: One subtree branch differs
	colors2 := []int{1, 2, 1, 1, 1, 1, 1}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors2)) // 3 (subtree [2,5,6] all color 1)

	// Test 3: All nodes different
	colors3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(maximumSubtreeOfSameColor(edges, colors3)) // 1 (each node alone)
}

func maximumSubtreeOfSameColor(edges [][]int, colors []int) int {
	n := len(colors)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := 0
	var dfs func(u, p int) (size int, same bool)
	dfs = func(u, p int) (size int, same bool) {
		size = 1
		same = true
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			childSize, childSame := dfs(v, u)
			if childSame && colors[v] == colors[u] {
				size += childSize
			} else {
				same = false
			}
		}
		if same && size > ans {
			ans = size
		}
		return
	}

	dfs(0, -1)
	return ans
}
