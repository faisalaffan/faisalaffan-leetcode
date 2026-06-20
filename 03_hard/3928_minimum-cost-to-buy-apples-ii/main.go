package main

// LeetCode #3928: Minimum Cost to Buy Apples II
// https://leetcode.com/problems/minimum-cost-to-buy-apples-ii/
// Difficulty: Hard [Paid]
//
// n stores, each with price for one apple. Roads connect stores
// with cost and tax. Buy one apple from each store. You can travel
// via roads, paying cost + tax. Find min total cost for each store
// (buy from that store or travel to another).
//
// Approach: Multi-source Dijkstra. Start from each store with
// initial cost = price[i]. Relax edges to find min cost for each
// store.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(minCost(3, []int{3, 5, 7}, [][]int{{0, 1, 2, 1}, {1, 2, 3, 2}}))
	// Example 2
	fmt.Println(minCost(2, []int{10, 20}, [][]int{{0, 1, 5, 0}}))
	// Edge: single store
	fmt.Println(minCost(1, []int{8}, [][]int{}))
}

type item struct {
	node int
	cost int
	idx  int
}

type minHeap []*item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *minHeap) Push(x interface{}) {
	n := len(*h)
	it := x.(*item)
	it.idx = n
	*h = append(*h, it)
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	it := old[n-1]
	it.idx = -1
	*h = old[0 : n-1]
	return it
}

func minCost(n int, prices []int, roads [][]int) []int {
	adj := make([][][3]int, n)
	for _, r := range roads {
		u, v, cost, tax := r[0], r[1], r[2], r[3]
		adj[u] = append(adj[u], [3]int{v, cost, tax})
		adj[v] = append(adj[v], [3]int{u, cost, tax})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = prices[i]
	}

	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, &item{node: i, cost: prices[i]})
	}

	for h.Len() > 0 {
		cur := heap.Pop(h).(*item)
		u, cu := cur.node, cur.cost
		if cu != dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v, cost, tax := edge[0], edge[1], edge[2]
			nd := cu + cost + tax
			if nd < dist[v] {
				dist[v] = nd
				heap.Push(h, &item{node: v, cost: nd})
			}
		}
	}

	return dist
}
