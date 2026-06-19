package main

// LeetCode #1786: Number of Restricted Paths From First to Last Node
// https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/
// Difficulty: Medium
// Time: O(E log V), Space: O(V + E)

import (
	"container/heap"
	"fmt"
)

const mod = 1_000_000_007

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func countRestrictedPaths(n int, edges [][]int) int {
	graph := make([][]Edge, n+1)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

	// Dijkstra from node n to all nodes
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = 1 << 60
	}
	dist[n] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, Item{n, 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		u := item.node
		if item.dist > dist[u] {
			continue
		}
		for _, e := range graph[u] {
			if nd := dist[u] + e.weight; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}

	// DP: count restricted paths
	memo := make([]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = -1
	}

	var dfs func(u int) int
	dfs = func(u int) int {
		if u == n {
			return 1
		}
		if memo[u] != -1 {
			return memo[u]
		}
		total := 0
		for _, e := range graph[u] {
			if dist[e.to] < dist[u] {
				total = (total + dfs(e.to)) % mod
			}
		}
		memo[u] = total
		return total
	}

	return dfs(1)
}

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}))
	// Expected: 3

	fmt.Println(countRestrictedPaths(7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}}))
	// Expected: 1

	fmt.Println(countRestrictedPaths(2, [][]int{{1, 2, 5}}))
	// Expected: 1
}
