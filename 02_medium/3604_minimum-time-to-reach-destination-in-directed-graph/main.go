package main

// LeetCode #3604: Minimum Time to Reach Destination in Directed Graph
// https://leetcode.com/problems/minimum-time-to-reach-destination-in-directed-graph/
// Difficulty: Medium
// Complexity: O((n+e) log n) time, O(n+e) space

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, time int
}

type Item struct {
	node, time int
	index       int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x any) { n := len(*pq); item := x.(*Item); item.index = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() any { old := *pq; n := len(old); item := old[n-1]; old[n-1] = nil; item.index = -1; *pq = old[:n-1]; return item }

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1, 5}, {0, 2, 2}, {1, 3, 1}, {2, 3, 3}}
	fmt.Println("Test 1:", MinimumTimeToReachDestinationInDirectedGraph(n, edges, 0, 3))
	// Test case 2
	n2 := 2
	edges2 := [][]int{{0, 1, 10}}
	fmt.Println("Test 2:", MinimumTimeToReachDestinationInDirectedGraph(n2, edges2, 0, 1))
	// Test case 3
	n3 := 3
	edges3 := [][]int{}
	fmt.Println("Test 3:", MinimumTimeToReachDestinationInDirectedGraph(n3, edges3, 0, 2))
}

func MinimumTimeToReachDestinationInDirectedGraph(n int, edges [][]int, start, end int) int {
	adj := make([][]Edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, time: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.node == end {
			return curr.time
		}
		if curr.time > dist[curr.node] {
			continue
		}
		for _, edge := range adj[curr.node] {
			nd := curr.time + edge.time
			if nd < dist[edge.to] {
				dist[edge.to] = nd
				heap.Push(pq, &Item{node: edge.to, time: nd})
			}
		}
	}
	return -1
}
