package main

import (
	"fmt"
)

// LeetCode #685: Redundant Connection II
// https://leetcode.com/problems/redundant-connection-ii/
// Difficulty: Hard
//
// Directed graph with n nodes and n edges. Either a node has in-degree 2
// (two parents) OR there is a cycle. Use Union-Find + in-degree tracking.

func main() {
	// [[1,2],[1,3],[2,3]] => [2,3]
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	// [[1,2],[2,3],[3,4],[4,1],[1,5]] => [4,1]
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
	// [[1,2],[2,3],[3,1]] => [3,1]
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 1}}))
	// [[2,1],[3,1],[1,4],[4,2]] => [3,1]
	fmt.Println(findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {1, 4}, {4, 2}}))
	// [[1,2],[2,3],[3,4],[1,4]] => [1,4]
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}}))
}

type uf struct {
	parent []int
}

func newUF(n int) *uf {
	p := make([]int, n+1)
	for i := range p {
		p[i] = i
	}
	return &uf{p}
}

func (u *uf) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *uf) union(x, y int) {
	rx, ry := u.find(x), u.find(y)
	if rx != ry {
		u.parent[ry] = rx
	}
}

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var cand1, cand2 []int

	// Phase 1: find node with in-degree 2
	for _, e := range edges {
		u, v := e[0], e[1]
		if parent[v] != v {
			cand1 = []int{parent[v], v}
			cand2 = []int{u, v}
			e[0], e[1] = -1, -1
		} else {
			parent[v] = u
		}
	}

	// Phase 2: Union-Find to detect cycle
	u := newUF(n)
	for _, e := range edges {
		if e[0] == -1 {
			continue
		}
		x, y := e[0], e[1]
		if u.find(x) == u.find(y) {
			if cand1 == nil {
				return e
			}
			return cand1
		}
		u.union(x, y)
	}

	return cand2
}
