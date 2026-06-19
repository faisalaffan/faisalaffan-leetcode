package main

// LeetCode #3778: Minimum Distance Excluding One Maximum Weighted Edge
// https://leetcode.com/problems/minimum-distance-excluding-one-maximum-weighted-edge/
// Difficulty: Medium [Paid]
// Time: O(E log V) | Space: O(V + E)

import (
	"container/heap"
	"fmt"
	"math"
)

type edge struct {
	to, w int
}

type state struct {
	node     int
	used     bool // whether max edge already excluded
	total    int
	maxEdge  int
}

type pqItem struct {
	total   int
	node    int
	used    bool
	maxEdge int
	index   int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].total-pq[i].maxEdge < pq[j].total-pq[j].maxEdge }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *priorityQueue) Push(x any) {
	item := x.(*pqItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[:n-1]
	return item
}

func minimumDistanceExcludingOneMaximumWeightedEdge(n int, edges [][]int) int {
	graph := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], edge{v, w})
		graph[v] = append(graph[v], edge{u, w})
	}

	// dist[node][used] = min effective cost (sum - maxEdge)
	dist := make([][2]int, n)
	for i := 0; i < n; i++ {
		dist[i][0] = math.MaxInt32
		dist[i][1] = math.MaxInt32
	}
	dist[0][0] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{node: 0, total: 0, maxEdge: 0, used: false})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*pqItem)
		effective := cur.total - cur.maxEdge
		if effective != dist[cur.node][btoi(cur.used)] {
			continue
		}
		if cur.node == n-1 {
			return effective
		}

		for _, e := range graph[cur.node] {
			// Option 1: Don't exclude this edge
			newTotal := cur.total + e.w
			newMax := cur.maxEdge
			if e.w > newMax {
				newMax = e.w
			}
			eff := newTotal - newMax
			if eff < dist[e.to][btoi(cur.used)] {
				dist[e.to][btoi(cur.used)] = eff
				heap.Push(pq, &pqItem{
					node: e.to, total: newTotal, maxEdge: newMax, used: cur.used,
				})
			}

			// Option 2: Exclude this edge (if not already excluded)
			if !cur.used {
				// Excluding means we don't add e.w to total, and track it as maxEdge
				// Actually: we just skip this edge's weight entirely
				if cur.total < dist[e.to][1] {
					dist[e.to][1] = cur.total
					heap.Push(pq, &pqItem{
						node: e.to, total: cur.total, maxEdge: cur.total, used: true,
					})
				}
			}
		}
	}
	return -1
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(4, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 1}, {0, 3, 10}}))
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(3, [][]int{{0, 1, 5}, {1, 2, 5}, {0, 2, 10}}))
	fmt.Println(minimumDistanceExcludingOneMaximumWeightedEdge(2, [][]int{{0, 1, 100}}))
}
