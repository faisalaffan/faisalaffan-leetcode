package main

// LeetCode #2473: Minimum Cost to Buy Apples
// https://leetcode.com/problems/minimum-cost-to-buy-apples/
// Difficulty: Medium
// Time: O(n * (n + m) log n) | Space: O(n + m)
// Run Dijkstra from each start node to find min round-trip cost for one apple.

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}

type Item struct {
	node, dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq; n := len(old); x := old[n-1]; *pq = old[:n-1]; return x
}

func main() {
	fmt.Println(minCost(5, [][]int{{1,2,2},{2,3,4},{3,4,6}}, []int{3,5,8,12,20}, 2, 1))
	fmt.Println(minCost(4, [][]int{{1,2,5},{2,3,1},{3,4,2}}, []int{10,10,10,10}, 3, 2))
}

func minCost(n int, roads [][]int, appleCost []int, k int, start int) []int64 {
	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, c := r[0]-1, r[1]-1, r[2]
		graph[u] = append(graph[u], Edge{v, c})
		graph[v] = append(graph[v], Edge{u, c})
	}

	ans := make([]int64, n)
	for s := 0; s < n; s++ {
		dist := dijkstra(graph, s, n)
		minCost := int64(math.MaxInt64)
		for i := 0; i < n; i++ {
			if dist[i] == math.MaxInt64 {
				continue
			}
			total := int64(appleCost[i]) + int64(dist[i])*(1+int64(k))
			if total < minCost {
				minCost = total
			}
		}
		ans[s] = minCost
	}
	return ans
}

func dijkstra(graph [][]Edge, src, n int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[src] = 0
	pq := &PriorityQueue{{src, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.dist + e.cost; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, Item{e.to, nd})
			}
		}
	}
	return dist
}
