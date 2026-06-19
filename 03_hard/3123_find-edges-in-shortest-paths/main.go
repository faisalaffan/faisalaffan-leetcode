package main

// LeetCode #3123: Find Edges in Shortest Paths
// https://leetcode.com/problems/find-edges-in-shortest-paths/
// Difficulty: Hard
//
// Given an undirected weighted graph, determine for each edge whether it
// belongs to at least one shortest path from node 0 to node n-1.
// Use two Dijkstras: distances from 0 and from n-1, then check each edge.

import (
	"container/heap"
	"fmt"
)

type Item struct {
	node, dist int
	idx        int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool   { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)        { pq[i], pq[j] = pq[j], pq[i]; pq[i].idx = i; pq[j].idx = j }
func (pq *PriorityQueue) Push(x interface{})  { it := x.(*Item); it.idx = len(*pq); *pq = append(*pq, it) }
func (pq *PriorityQueue) Pop() interface{}    { old := *pq; n := len(old); it := old[n-1]; it.idx = -1; *pq = old[:n-1]; return it }

func dijkstra(n int, adj [][][]int, start int) []int {
	const INF = 1 << 60
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[start] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, dist: 0})
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u, d := item.node, item.dist
		if d > dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v, w := edge[0], edge[1]
			if nd := d + w; nd < dist[v] {
				dist[v] = nd
				heap.Push(pq, &Item{node: v, dist: nd})
			}
		}
	}
	return dist
}

func findEdgesInShortestPaths(n int, edges [][]int) []bool {
	adj := make([][][]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], []int{v, w})
		adj[v] = append(adj[v], []int{u, w})
	}

	distFromStart := dijkstra(n, adj, 0)
	distFromEnd := dijkstra(n, adj, n-1)
	shortest := distFromStart[n-1]

	ans := make([]bool, len(edges))
	for i, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if distFromStart[u]+w+distFromEnd[v] == shortest ||
			distFromStart[v]+w+distFromEnd[u] == shortest {
			ans[i] = true
		}
	}
	return ans
}

func main() {
	// Test case 1
	n := 6
	edges := [][]int{
		{0, 1, 4}, {0, 2, 1}, {1, 3, 2}, {1, 4, 3},
		{1, 5, 1}, {2, 3, 1}, {3, 5, 3}, {4, 5, 2},
	}
	fmt.Println("Test 1:", findEdgesInShortestPaths(n, edges))
	// Expected: [true, false, true, false, true, false, true, false]

	// Test case 2: simple 2-node
	n2 := 2
	edges2 := [][]int{{0, 1, 5}}
	fmt.Println("Test 2:", findEdgesInShortestPaths(n2, edges2))
	// Expected: [true]

	// Test case 3: no path
	n3 := 3
	edges3 := [][]int{{0, 1, 1}}
	fmt.Println("Test 3:", findEdgesInShortestPaths(n3, edges3))
	// Expected: [false] (node 2 unreachable from 0)
}
