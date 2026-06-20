# 2093 — Minimum Cost To Reach City With Discounts

## Deskripsi

**Soal:** [2093. Minimum Cost To Reach City With Discounts](https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((n+m) * discounts * log(n*discounts))  
**Kompleksitas Ruang:** O(n * discounts)

**Algoritma:** Queue (antrian FIFO), Heap (priority queue), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func minimumCost(n int, highways [][]int, discounts int) int`

## Solusi Go

```go
package main

// LeetCode #2093: Minimum Cost to Reach City With Discounts
// https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/
// Difficulty: Medium [Paid]
// Time: O((n+m) * discounts * log(n*discounts)) | Space: O(n * discounts)

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}

type State struct {
	city    int
	discounts int
	cost    int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func minimumCost(n int, highways [][]int, discounts int) int {
	// Build adjacency list
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]Edge, n)
	for _, h := range highways {
		u, v, c := h[0], h[1], h[2]
		adj[u] = append(adj[u], Edge{v, c})
		adj[v] = append(adj[v], Edge{u, c})
	}

	// dist[city][discountsUsed] = min cost
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int, discounts+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, State{0, 0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		if cur.cost > dist[cur.city][cur.discounts] {
			continue
		}

		for _, e := range adj[cur.city] {
			// Without discount
			nc := cur.cost + e.cost
			if nc < dist[e.to][cur.discounts] {
				dist[e.to][cur.discounts] = nc
				heap.Push(pq, State{e.to, cur.discounts, nc})
			}
			// With discount
			if cur.discounts < discounts {
				nc2 := cur.cost + e.cost/2
				if nc2 < dist[e.to][cur.discounts+1] {
					dist[e.to][cur.discounts+1] = nc2
					heap.Push(pq, State{e.to, cur.discounts + 1, nc2})
				}
			}
		}
	}

	result := math.MaxInt32
	for d := 0; d <= discounts; d++ {
		if dist[n-1][d] < result {
			result = dist[n-1][d]
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func main() {
	// Test case 1
	n1 := 5
	highways1 := [][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 3}, {1, 4, 1}, {3, 4, 5}}
	discounts1 := 1
	fmt.Println("Test 1:", minimumCost(n1, highways1, discounts1))
	// Expected: 6

	// Test case 2
	n2 := 4
	highways2 := [][]int{{1, 3, 17}, {1, 2, 7}, {3, 2, 5}, {0, 1, 6}, {3, 0, 20}}
	discounts2 := 20
	fmt.Println("Test 2:", minimumCost(n2, highways2, discounts2))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", minimumCost(2, [][]int{{0, 1, 10}}, 0))
	// Expected: 10
}
```
