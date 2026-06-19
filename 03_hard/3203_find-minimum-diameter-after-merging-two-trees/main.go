package main

// LeetCode #3203: Find Minimum Diameter After Merging Two Trees
// https://leetcode.com/problems/find-minimum-diameter-after-merging-two-trees/
// Difficulty: Hard
//
// Given two trees (undirected acyclic graphs), we connect one node from each
// tree with an edge, forming a new tree. Find the minimum possible diameter
// of the resulting tree. The answer is max(d1, d2, ceil(d1/2)+ceil(d2/2)+1).

import (
	"fmt"
)

func main() {
	edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	edges2 := [][]int{{0, 1}}
	fmt.Println(minimumDiameterAfterMergingTwoTrees(edges1, edges2))
}

func minimumDiameterAfterMergingTwoTrees(edges1, edges2 [][]int) int {
	d1 := treeDiameter(edges1)
	d2 := treeDiameter(edges2)
	// Minimum diameter after merging
	merge := (d1+1)/2 + (d2+1)/2 + 1
	ans := d1
	if d2 > ans {
		ans = d2
	}
	if merge > ans {
		ans = merge
	}
	return ans
}

func treeDiameter(edges [][]int) int {
	n := len(edges) + 1
	if n <= 1 {
		return 0
	}
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS from 0 to find farthest node
	far1, _ := bfs(adj, 0)
	// BFS from farthest node to get diameter
	_, dist := bfs(adj, far1)
	return dist
}

func bfs(adj [][]int, start int) (farthest, dist int) {
	n := len(adj)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true
	distArr := make([]int, n)

	farthest = start
	maxDist := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true
				distArr[v] = distArr[u] + 1
				queue = append(queue, v)
				if distArr[v] > maxDist {
					maxDist = distArr[v]
					farthest = v
				}
			}
		}
	}
	return farthest, maxDist
}
