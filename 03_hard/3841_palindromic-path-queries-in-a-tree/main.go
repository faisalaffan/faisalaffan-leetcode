package main

// LeetCode #3841: Palindromic Path Queries in a Tree
// https://leetcode.com/problems/palindromic-path-queries-in-a-tree/
// Difficulty: Hard
//
// For each query (u, v), determine if nodes on path u-v can be
// rearranged into a palindrome. A set of chars can form a palindrome
// iff at most one char has odd frequency.
//
// Approach: Assign each char a 26-bit mask. Compute prefix XOR from
// root. Path XOR = pref[u] ^ pref[v] ^ mask[lca]. Palindromic iff
// popcount(path XOR) <= 1. Use binary lifting for LCA.

import (
	"fmt"
	"math/bits"
)

func main() {
	// Example 1
	n1 := 5
	edges1 := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}
	s1 := "abcba"
	q1 := []string{"0,4", "1,3", "2,3"}
	fmt.Println(palindromicPathQueries(n1, edges1, s1, q1))

	// Example 2: single node
	n2 := 1
	edges2 := [][]int{}
	s2 := "a"
	q2 := []string{"0,0"}
	fmt.Println(palindromicPathQueries(n2, edges2, s2, q2))

	// Edge: two nodes same char
	n3 := 2
	edges3 := [][]int{{0, 1}}
	s3 := "aa"
	q3 := []string{"0,1"}
	fmt.Println(palindromicPathQueries(n3, edges3, s3, q3))
}

func palindromicPathQueries(n int, edges [][]int, s string, queries []string) []bool {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	LOG := 0
	for (1 << LOG) <= n {
		LOG++
	}
	up := make([][]int, n)
	for i := range up {
		up[i] = make([]int, LOG)
		for j := range up[i] {
			up[i][j] = -1
		}
	}
	depth := make([]int, n)
	pref := make([]int, n)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		up[u][0] = p
		for j := 1; j < LOG; j++ {
			if up[u][j-1] != -1 {
				up[u][j] = up[up[u][j-1]][j-1]
			}
		}
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			depth[v] = depth[u] + 1
			pref[v] = pref[u] ^ (1 << (s[v] - 'a'))
			dfs(v, u)
		}
	}
	pref[0] = 1 << (s[0] - 'a')
	dfs(0, -1)

	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for j := LOG - 1; j >= 0; j-- {
			if up[u][j] != -1 && depth[up[u][j]] >= depth[v] {
				u = up[u][j]
			}
		}
		if u == v {
			return u
		}
		for j := LOG - 1; j >= 0; j-- {
			if up[u][j] != up[v][j] {
				u = up[u][j]
				v = up[v][j]
			}
		}
		return up[u][0]
	}

	ans := make([]bool, len(queries))
	for idx, q := range queries {
		var u, v int
		fmt.Sscanf(q, "%d,%d", &u, &v)
		w := lca(u, v)
		xor := pref[u] ^ pref[v] ^ (1 << (s[w] - 'a'))
		ans[idx] = bits.OnesCount(uint(xor)) <= 1
	}
	return ans
}
