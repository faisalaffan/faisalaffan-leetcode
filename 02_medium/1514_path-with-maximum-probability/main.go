package main

// LeetCode #1514: Path with Maximum Probability
// https://leetcode.com/problems/path-with-maximum-probability/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(MaxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println(MaxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	fmt.Println(MaxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}

type Edge struct {
	to   int
	prob float64
}

type Item struct {
	node   int
	prob   float64
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].prob > pq[j].prob } // max-heap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Time: O(E log V), Space: O(E + V)
	// Build adjacency list
	graph := make([][]Edge, n)
	for i, e := range edges {
		u, v := e[0], e[1]
		p := succProb[i]
		graph[u] = append(graph[u], Edge{v, p})
		graph[v] = append(graph[v], Edge{u, p})
	}

	// Dijkstra-like (max probability)
	prob := make([]float64, n)
	prob[start] = 1.0

	pq := &PriorityQueue{}
	heap.Push(pq, &Item{node: start, prob: 1.0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.node
		curProb := item.prob

		if node == end {
			return curProb
		}

		if curProb < prob[node] {
			continue
		}

		for _, edge := range graph[node] {
			newProb := curProb * edge.prob
			if newProb > prob[edge.to] {
				prob[edge.to] = newProb
				heap.Push(pq, &Item{node: edge.to, prob: newProb})
			}
		}
	}

	return 0.0
}
