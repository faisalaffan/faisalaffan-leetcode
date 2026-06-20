package main

// LeetCode #3924: Minimum Threshold Path With Limited Heavy Edges
// https://leetcode.com/problems/minimum-threshold-path-with-limited-heavy-edges/
// Difficulty: Hard
//
// Given an undirected weighted graph, find a path from node 0 to
// node n-1 that minimizes the number of "heavy" edges (edges with
// weight > threshold). If multiple paths have the same count,
// choose the one with minimum total weight.
//
// Approach: Dijkstra-like with state (node, heavyCount, totalWeight).
// Priority queue orders by (heavyCount, totalWeight).

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minThresholdPath(4, [][]int{{0, 1, 5}, {1, 2, 3}, {2, 3, 1}, {0, 3, 10}}, 4))
	// Example 2
	fmt.Println(minThresholdPath(3, [][]int{{0, 1, 2}, {1, 2, 8}}, 5))
	// Edge: single node
	fmt.Println(minThresholdPath(1, [][]int{}, 1))
}

type state struct {
	node       int
	heavyCount int
	totalDist  int
	index      int
}

type priorityQueue []*state

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].heavyCount != pq[j].heavyCount {
		return pq[i].heavyCount < pq[j].heavyCount
	}
	return pq[i].totalDist < pq[j].totalDist
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*state)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func minThresholdPath(n int, edges [][]int, threshold int) int {
	if n <= 1 {
		return 0
	}

	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// dist[node][heavy] = min total distance
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &state{node: 0, heavyCount: 0, totalDist: 0})
	dist[0][0] = 0

	for pq.Len() > 0 {
		s := heap.Pop(pq).(*state)
		if s.node == n-1 {
			return s.totalDist
		}
		if s.totalDist > dist[s.node][s.heavyCount] {
			continue
		}
		for _, edge := range adj[s.node] {
			v, w := edge[0], edge[1]
			nh := s.heavyCount
			if w > threshold {
				nh++
			}
			nd := s.totalDist + w
			if nh <= n && nd < dist[v][nh] {
				dist[v][nh] = nd
				heap.Push(pq, &state{node: v, heavyCount: nh, totalDist: nd})
			}
		}
	}

	return -1
}
