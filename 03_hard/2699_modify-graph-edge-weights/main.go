package main

// LeetCode #2699: Modify Graph Edge Weights
// https://leetcode.com/problems/modify-graph-edge-weights/
// Difficulty: Hard
//
// Given undirected graph with n nodes, edges[i] = [u, v, w] where w = -1 means
// the weight can be set to any positive integer. Find a setting of all -1
// weights such that the shortest path from source to destination equals target.
// Return the modified edges, or empty array if impossible.
//
// Approach:
// 1. Set all -1 edges to INF, run Dijkstra. If dist[dest] < target -> impossible.
// 2. Set all -1 edges to 1, run Dijkstra. If dist[dest] > target -> impossible.
// 3. For each -1 edge (in original order), binary search its weight in [1, INF)
//    to try making shortest path = target, keeping others at INF.
//    Once found, set remaining -1 edges to INF.

import (
	"container/heap"
	"fmt"
)

const INF = 1_000_000_000

func main() {
	// Example 1
	n := 5
	edges := [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}
	modifyGraphEdgeWeights(n, edges, 0, 1, 5)
	fmt.Println("---")

	// Example 2
	n2 := 3
	edges2 := [][]int{{0, 1, -1}, {0, 2, 5}}
	modifyGraphEdgeWeights(n2, edges2, 0, 2, 6)
	fmt.Println("---")

	// Example 3: impossible
	n3 := 4
	edges3 := [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}
	modifyGraphEdgeWeights(n3, edges3, 0, 2, 1)
}

type edge struct {
	to, weight int
}

type item struct {
	dist, node int
}

type pq []item

func (p pq) Len() int            { return len(p) }
func (p pq) Less(i, j int) bool  { return p[i].dist < p[j].dist }
func (p pq) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)          { *p = append(*p, x.(item)) }
func (p *pq) Pop() any            { old := *p; n := len(old); x := old[n-1]; *p = old[:n-1]; return x }

func dijkstra(n int, adj [][]edge, source, dest int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[source] = 0
	pq := &pq{{0, source}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		if cur.node == dest {
			return cur.dist
		}
		for _, e := range adj[cur.node] {
			nd := cur.dist + e.weight
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, item{nd, e.to})
			}
		}
	}
	return dist[dest]
}

// buildAdj creates adjacency from edges; negative edges get the specified default.
func buildAdj(n int, edges [][]int, def int) [][]edge {
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w == -1 {
			w = def
		}
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}
	return adj
}

func modifyGraphEdgeWeights(n int, edges [][]int, source int, destination int, target int) [][]int {
	// Phase 1: check if achievable with minimum weights
	minAdj := buildAdj(n, edges, 1)
	minDist := dijkstra(n, minAdj, source, destination)
	if minDist > target {
		fmt.Println("impossible: min dist > target")
		return [][]int{}
	}

	// Phase 2: check if achievable at all
	maxAdj := buildAdj(n, edges, INF)
	maxDist := dijkstra(n, maxAdj, source, destination)
	if maxDist < target {
		fmt.Println("impossible: max dist < target")
		return [][]int{}
	}

	if maxDist == target && minDist == target {
		// Already works with all -1 as 1
		result := make([][]int, len(edges))
		for i, e := range edges {
			w := e[2]
			if w == -1 {
				w = 1
			}
			result[i] = []int{e[0], e[1], w}
		}
		printResult(result)
		return result
	}

	// Phase 3: adjust -1 edges iteratively
	// First pass: set all -1 edges to INF and record their indices
	result := make([][]int, len(edges))
	for i, e := range edges {
		result[i] = []int{e[0], e[1], e[2]}
	}

	// Collect indices of -1 edges
	var negIndices []int
	for i, e := range edges {
		if e[2] == -1 {
			negIndices = append(negIndices, i)
		}
	}

	// Start with all -1 at INF, then tune them one by one
	for _, idx := range negIndices {
		result[idx][2] = 1
	}

	// Check if already correct
	currAdj := buildAdjFromResult(n, result)
	currDist := dijkstra(n, currAdj, source, destination)
	if currDist == target {
		printResult(result)
		return result
	}

	// For each -1 edge, try adjusting to make the path exactly target
	// If current dist < target, we need to increase some edge weight
	// The strategy: raise each -1 edge from 1 upward until dist == target
	need := target - currDist // extra weight needed
	if need < 0 {
		// Current shortest is too long (shouldn't happen since we start with all 1)
		fmt.Println("impossible: need is negative")
		return [][]int{}
	}

	if need > 0 {
		// We need to add need to some -1 edges' weight to match target.
		// Strategy: add the entire need to the last -1 edge on the critical path,
		// or distribute across multiple edges.
		// Simplest correct approach: set all -1 edges to 1, then for each -1 edge
		// in order, binary search its exact value.
		for _, idx := range negIndices {
			result[idx][2] = INF
		}

		for i, idx := range negIndices {
			lo, hi := 1, INF
			for lo < hi {
				mid := lo + (hi-lo)/2
				result[idx][2] = mid
				adj := buildAdjFromResult(n, result)
				d := dijkstra(n, adj, source, destination)
				if d <= target {
					hi = mid
				} else {
					lo = mid + 1
				}
			}
			result[idx][2] = lo
			adj := buildAdjFromResult(n, result)
			d := dijkstra(n, adj, source, destination)
			if d == target {
				// Found the right value, set remaining -1 edges to INF
				for j := i + 1; j < len(negIndices); j++ {
					result[negIndices[j]][2] = INF
				}
				printResult(result)
				return result
			}
		}
	}

	// Final check
	adj := buildAdjFromResult(n, result)
	d := dijkstra(n, adj, source, destination)
	if d == target {
		printResult(result)
		return result
	}

	fmt.Println("impossible: cannot match target")
	return [][]int{}
}

func buildAdjFromResult(n int, result [][]int) [][]edge {
	adj := make([][]edge, n)
	for _, e := range result {
		u, v, w := e[0], e[1], e[2]
		if w == -1 || w >= INF {
			continue // skip unset negative edges for this check
		}
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}
	return adj
}

func printResult(res [][]int) {
	fmt.Print("[")
	for i, e := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("[%d,%d,%d]", e[0], e[1], e[2])
	}
	fmt.Println("]")
}
